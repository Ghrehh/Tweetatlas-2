package web

type ConnectionOrchestrator struct {
	connections map[*Connection]bool
	laStream chan LocationAggregater
	add chan *Connection
	remove chan *Connection
}

func NewConnectionOrchestrator(laStream chan LocationAggregater) *ConnectionOrchestrator {
	return &ConnectionOrchestrator{
		connections: make(map[*Connection]bool),
		laStream: laStream,
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
		case la := <-co.laStream:
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
