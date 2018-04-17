package main

import (
	"net/http"
	"log"

	"github.com/ghrehh/tweetatlas/twitterstream"
	"github.com/ghrehh/tweetatlas/utils"
	"github.com/ghrehh/tweetatlas/web"

	"github.com/ghrehh/findlocation"
	"github.com/dghubble/go-twitter/twitter"
	"github.com/gorilla/handlers"
  "github.com/gorilla/mux"
	"github.com/gorilla/websocket"
)

func main() {
	log.Print("App Started")

	upgrader := websocket.Upgrader{
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}

	// get the search phrase(s)
	filter := utils.GetStreamFilter()

	// location aggregate stores our location results
	locationAggregate := web.NewLocationAggregate(filter)

	// location finder is a package that attempts to find a Twitter user's location.
	locationFinder := findlocation.NewLocationFinder()

	// connection orchestartor manages the websocket connections
	co := web.NewConnectionOrchestrator()

	twitterKeys := utils.GetOauthKeys()
	streamer := twitterstream.NewTweetStream(twitterKeys)
	stream, err := twitterstream.FilterStream(&streamer, filter)

	if err != nil {
		log.Print(err)
	}

	// Handle different types returned by the stream
	demux := twitter.NewSwitchDemux()
	demux.Tweet = func(tweet *twitter.Tweet) {
		location := locationFinder.FindLocation(tweet.User.Location)
		locationAggregate.AddParsedLocation(location)
		co.LaStream <- locationAggregate
	}

	demux.Other = func(message interface{}) {
		log.Print(message)
	}

	go demux.HandleChan(stream.Messages)

	// Start the connection orchestrator
	go co.Run()

	// Routing
	r := mux.NewRouter()
	r.Headers("Content-Type", "application/json")
	r.HandleFunc("/tweets", func(w http.ResponseWriter, r *http.Request) {
		web.ServeWebsocket(co, w, r, &upgrader)
	})

	http.ListenAndServe(utils.GetPort(), handlers.CORS()(r))
}
