package twitterstream

import (
	"github.com/ghrehh/tweetatlas/keys"
	"github.com/dghubble/go-twitter/twitter"
	"github.com/dghubble/oauth1"
)


type tweetStreamer interface {
	FilterStream(*twitter.StreamFilterParams) (*twitter.Stream, error)
}

type TweetStreamer struct {
	Tc *twitter.Client
}

func (t *TweetStreamer) FilterStream(params *twitter.StreamFilterParams) (*twitter.Stream, error) {
	return t.Tc.Streams.Filter(params)
}

func NewTweetStreamer(k keys.OauthKeys) TweetStreamer {
	config := oauth1.NewConfig(k.ConsumerKey, k.ConsumerSecret)
	token := oauth1.NewToken(k.AccessToken, k.AccessSecret)
	httpClient := config.Client(oauth1.NoContext, token)
	client := twitter.NewClient(httpClient)

	return TweetStreamer{client}
}

func FilterStream(ts tweetStreamer, filter []string) (*twitter.Stream, error)  {
	params := &twitter.StreamFilterParams{
    Track: filter,
    StallWarnings: twitter.Bool(true),
	}

	return ts.FilterStream(params)
}
