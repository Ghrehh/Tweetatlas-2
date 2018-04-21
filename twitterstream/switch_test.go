package twitterstream

import (
	"testing"
)

func TestSwitchRun(t *testing.T) {
	handlerCalls := 0
	stream := make(chan interface{}, 2)

	s := Switch{
		stream,
		func(interface{}) {
			handlerCalls++
		},
	}

	stream <- "foo"
	stream <- "bar"
	close(stream)

	s.Run()

	if handlerCalls != 2 {
		t.Error("expected handler to be called twice")
	}
}

func TestSwitchSwitchStream(t *testing.T) {
	stream := make(chan interface{})
	newStream := make(chan interface{})

	s := Switch{
		stream,
		func(interface{}) {},
	}

	s.SwitchStream(newStream)

	if <-stream != nil {
		t.Error("expected original stream to be closed")
	}

	if s.Stream != newStream {
		t.Error("expected StreamSwitch Stream to be newStream")
	}
}
