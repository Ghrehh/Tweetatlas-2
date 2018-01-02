package twitterstream

import (
	"github.com/dghubble/go-twitter/twitter"
	"testing"
)

type fakeTweetStreamer struct {
	tc *twitter.Client
}

func (f fakeTweetStreamer) FilterStream(params *twitter.StreamFilterParams) (*twitter.Stream, error) {
	return &twitter.Stream{}, nil
}

func TestFilterStream(t *testing.T) {
	filter := []string{"happy"}
	streamer := fakeTweetStreamer{&twitter.Client{}}

	_, err := FilterStream(streamer, filter)

	if err != nil {
		t.Error("expected no error to get raised")
	}
}
