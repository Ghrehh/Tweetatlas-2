package twitterstream

import (
	"github.com/dghubble/go-twitter/twitter"
	"testing"
	"errors"
	"reflect"
)

type fakeStreams struct {}

func (fs fakeStreams) Filter(s *twitter.StreamFilterParams) (*twitter.Stream, error) {
	if reflect.DeepEqual(s.Track, []string{"foo"}) {
		return &twitter.Stream{}, nil
	}

	return &twitter.Stream{}, errors.New("error")
}

func TestFilterStream(t *testing.T) {
	sh := StreamHandler{fakeStreams{}}

	_, err := sh.CreateFilteredStream([]string{"foo"})

	if err != nil {
		t.Error("expected filter params to be passed correctly")
	}
}
