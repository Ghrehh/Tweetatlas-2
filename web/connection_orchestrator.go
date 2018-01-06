package web

type ConnectionOrchestrator struct {
	connections map[*Connection]bool
	LaStream chan LocationAggregater
	add chan *Connection
	remove chan *Connection
}

func NewConnectionOrchestrator() *ConnectionOrchestrator {
	return &ConnectionOrchestrator{
		connections: make(map[*Connection]bool),
		LaStream: make(chan LocationAggregater),
		add: make(chan *Connection),
		remove: make(chan *Connection),
	}
}

func (co *ConnectionOrchestrator) Run() {
	for {
		select {
		case connection := <- co.add:
			co.connections[connection] = true
		case connection := <- co.remove:
			if _, ok := co.connections[connection]; ok {
				delete(co.connections, connection)
				close(connection.dataStream)
			}
		case la := <- co.LaStream:
			laJSON := la.ToJSON()

			for connection := range co.connections {
				select {
				case connection.dataStream <- &laJSON:
				default:
					close(connection.dataStream)
					delete(co.connections, connection)
				}
			}
		}
	}
}
