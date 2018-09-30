package twitterstream

import (
	"log"
	"time"
)

type Scheduler struct{
	Switch Switcher
	StreamHandler CreateFilteredStreamer
	Filters []string
	FilterIndex int
	SearchDuration int
	LocationAggregate LocationAggregater
}

func (s *Scheduler) cycleFilter() {
	if s.FilterIndex == len(s.Filters) - 1 {
		s.FilterIndex = 0
		return
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

	nextSearch := time.Now().Add(time.Duration(s.SearchDuration) * time.Second)

	s.Switch.SwitchStream(stream)
	s.LocationAggregate.Reset(s.FilterIndex, nextSearch)
	s.Switch.Run()
}

func (s Scheduler) Run() {
	go s.StartStream()

	for {
		time.Sleep(time.Duration(s.SearchDuration) * time.Second)
		s.cycleFilter()
		go s.StartStream()
	}
}
