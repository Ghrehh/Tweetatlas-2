package twitterstream

import (
	"github.com/dghubble/go-twitter/twitter"
	"github.com/dghubble/oauth1"
	"testing"
	"reflect"
)

type fakeTweetStreamer struct {
	tc *twitter.Client
}

func (f fakeTweetStreamer) FilterStream(params *twitter.StreamFilterParams) (*twitter.Stream, error) {
	return &twitter.Stream{}, nil
}

func TestNewOauthConfig(t *testing.T) {
	config, token := NewOauthConfig(OauthKeys{})

	if reflect.TypeOf(config) != reflect.TypeOf(&oauth1.Config{}) {
		t.Error("expected the correct type")
	}

	if reflect.TypeOf(token) != reflect.TypeOf(&oauth1.Token{}) {
		t.Error("expected the correct type")
	}
}

func TestNewTwitterClient(t *testing.T) {
	client := NewTwitterClient(&oauth1.Config{}, &oauth1.Token{})

	if reflect.TypeOf(client) != reflect.TypeOf(&twitter.Client{}) {
		t.Error("expected the correct type")
	}
}

func TestFilterStream(t *testing.T) {
	filter := []string{"happy"}
	streamer := fakeTweetStreamer{&twitter.Client{}}

	stream, err := FilterStream(streamer, filter)

	if reflect.TypeOf(stream) != reflect.TypeOf(&twitter.Stream{}) {
		t.Error("expected the correct type")
	}

	if err != nil {
		t.Error("expected no error to get raised")
	}
}
