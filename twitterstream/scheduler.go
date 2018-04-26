package twitterstream

import (
	"log"
	"time"
	"fmt"
)

type Scheduler struct{
	Switch Switcher
	StreamHandler CreateFilteredStreamer
	Filters []string
	FilterIndex int
	SearchDuration int
}

func (s *Scheduler) incrementFilterIndex() {
	if s.FilterIndex == len(s.Filters) - 1 {
		s.FilterIndex = 0
	}

	s.FilterIndex ++
}

func (s Scheduler) StartStream () {
	filter := s.Filters[s.FilterIndex]
	stream, err := s.StreamHandler.CreateFilteredStream(
		[]string{filter},
	)

	if err != nil {
		log.Print(err)
	}

	s.Switch.SwitchStream(stream)
	s.Switch.Run()
}

func (s Scheduler) Run() {
	go s.StartStream()

	for {
		time.Sleep(time.Duration(s.SearchDuration) * time.Second)
		fmt.Println(s.Filters[s.FilterIndex])
		s.incrementFilterIndex()
		go s.StartStream()
	}
}
