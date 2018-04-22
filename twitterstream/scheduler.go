package twitterstream

import (
	"log"
)

type Scheduler struct{
	Switch Switcher
	StreamHandler CreateFilteredStreamer
	Filters []string
	FilterIndex int
	SearchDuration int
	LastSearchBegan int
}

func (s *Scheduler) Run() {
	stream, err := s.StreamHandler.CreateFilteredStream(
		s.Filters,
	)

	if err != nil {
		log.Print(err)
	}

	s.Switch.SwitchStream(stream)
	s.Switch.Run()
}
