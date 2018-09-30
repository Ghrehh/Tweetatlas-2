package twitterstream

import (
	"testing"
)

type fakeStreamWrapper struct {
	Messages chan interface{}
	StopFunction func()
}

func (f fakeStreamWrapper) Stop() {
	f.StopFunction()
}

func (f fakeStreamWrapper) GetMessages() chan interface{} {
	return f.Messages
}

func TestSwitchRun(t *testing.T) {
	handlerCalls := 0
	stream := fakeStreamWrapper{make(chan interface{}, 2), func(){}}

	s := Switch{
		stream,
		func(interface{}) {
			handlerCalls++
		},
	}

	stream.GetMessages() <- "foo"
	stream.GetMessages() <- "bar"
	close(stream.GetMessages())

	s.Run()

	if handlerCalls != 2 {
		t.Error("expected handler to be called twice")
	}
}

func TestSwitchSwitchStream(t *testing.T) {
	stopCalls := 0

	stream1 := fakeStreamWrapper{
		make(chan interface{}),
		func() {
			stopCalls++
		},
	}

	stream2 := fakeStreamWrapper{make(chan interface{}), func(){}}

	s := Switch{stream1, func(interface{}){}}
	s.SwitchStream(stream2)

	if stopCalls != 1 {
		t.Error("expected stop to be called once")
	}

	if s.Stream.GetMessages() != stream2.GetMessages() {
		t.Error("expected StreamSwitch Stream to be stream2")
	}
}
