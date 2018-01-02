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

func NewOauthConfig(k keys.OauthKeys) (*oauth1.Config, *oauth1.Token) {
	config := oauth1.NewConfig(k.ConsumerKey, k.ConsumerSecret)
	token := oauth1.NewToken(k.AccessToken, k.AccessSecret)

	return config, token
}

func NewTwitterClient(c *oauth1.Config, t *oauth1.Token) *twitter.Client {
	httpClient := c.Client(oauth1.NoContext, t)

	return twitter.NewClient(httpClient)
}

func FilterStream(ts tweetStreamer, filter []string) (*twitter.Stream, error)  {
	params := &twitter.StreamFilterParams{
    Track: filter,
    StallWarnings: twitter.Bool(true),
	}

	return ts.FilterStream(params)
}
