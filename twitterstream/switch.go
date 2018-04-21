package twitterstream

type Switch struct {
	Stream chan interface{}
	Handler func(interface{})
}

func (s Switch) Run() {
	for message := range s.Stream {
		if message == nil {
			break
		}

		s.Handler(message)
	}
}

func (s *Switch) SwitchStream(stream chan interface{}) {
	close(s.Stream)
	s.Stream = stream
}
