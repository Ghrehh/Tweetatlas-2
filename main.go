package main

import (
	"net/http"
	"os"
	"strings"

	"github.com/ghrehh/tweetatlas/twitterstream"
	"github.com/ghrehh/tweetatlas/keys"
	"github.com/ghrehh/tweetatlas/web"

	"github.com/dghubble/go-twitter/twitter"
	"github.com/gorilla/handlers"
  "github.com/gorilla/mux"
)

func main() {
	tweetStream := make(chan *twitter.Tweet)

	twitterKeys := keys.Parse(keys.Load())
	streamer := twitterstream.NewTweetStreamer(twitterKeys)
	filter := []string{"happy"}
	stream, _ := twitterstream.FilterStream(&streamer, filter)

	demux := twitter.NewSwitchDemux()
	demux.Tweet = func(tweet *twitter.Tweet) {
		tweetStream <- tweet
	}

	go demux.HandleChan(stream.Messages)

	co := web.NewConnectionOrchestrator(tweetStream)
	go co.Run()

	r := mux.NewRouter()
	r.Headers("Content-Type", "application/json")
	r.HandleFunc("/tweets", func(w http.ResponseWriter, r *http.Request) {
		web.ServeWebsocket(co, w, r)
	})

	userSpecifiedPort := os.Getenv("PORT")
	port := "5555"

	if userSpecifiedPort != "" {
		port = userSpecifiedPort
	}

	http.ListenAndServe(strings.Join([]string{":", port}, ""), handlers.CORS()(r))
}
