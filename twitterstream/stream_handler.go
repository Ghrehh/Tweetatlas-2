package twitterstream

import (
	"github.com/ghrehh/tweetatlas/utils"
	"github.com/dghubble/go-twitter/twitter"
	"github.com/dghubble/oauth1"
)

// interface for the streamWrapper type
type StopperGetMessages interface {
	Stop()
	GetMessages() chan interface{}
}

// wraps the go-twitter Stream type to give us a Messages getter
type streamWrapper struct {
	Stream *twitter.Stream
}

func (s streamWrapper) Stop() {
	s.Stream.Stop()
}

func (s streamWrapper) GetMessages() chan interface{} {
	return s.Stream.Messages
}

// interface for the go-twitter StreamService type
type filterer interface {
	Filter(*twitter.StreamFilterParams) (*twitter.Stream, error)
}

// interface for our StreamHandler type
type CreateFilteredStreamer interface {
	CreateFilteredStream([]string) (StopperGetMessages, error)
}

// A StreamHandler wraps the go-twitter StreamService type to allow us
// to easily stop and start new filtered streams.
type StreamHandler struct {
	 streams filterer
}

func (cs StreamHandler) CreateFilteredStream(filter []string) (StopperGetMessages, error) {
	params := &twitter.StreamFilterParams{
    Track: filter,
	}

	stream, err := cs.streams.Filter(params)
	wrappedStream := streamWrapper{stream}

	return wrappedStream, err
}

// helper method for initialising the streamService type wrapped by our StreamHandler
func NewStreamHandler(k utils.OauthKeys) StreamHandler {
	config := oauth1.NewConfig(k.ConsumerKey, k.ConsumerSecret)
	token := oauth1.NewToken(k.AccessToken, k.AccessSecret)
	httpClient := config.Client(oauth1.NoContext, token)
	streamService := twitter.NewClient(httpClient).Streams

	return StreamHandler{streamService}
}
