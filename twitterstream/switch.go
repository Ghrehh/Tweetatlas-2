package twitterstream

type Switcher interface {
	Run()
	SwitchStream(StopperGetMessages)
}

type Switch struct {
	Stream StopperGetMessages
	Handler func(interface{})
}

func (s Switch) Run() {
	for message := range s.Stream.GetMessages() {
		if message == nil {
			break
		}

		s.Handler(message)
	}
}

func (s *Switch) SwitchStream(stream StopperGetMessages) {
	if s.Stream != nil {
		s.Stream.Stop()
		close(s.Stream.GetMessages())
	}

	s.Stream = stream
}
