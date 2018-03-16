# Tweetatlas Server

## Twitter Keys

To run this app, you must have a set of Twitter API keys. Please refer to the
[Twitter docs](https://developer.twitter. com/en/docs) for information on how to
acquire those.

To make running the app locally easier, you can copy the
`config/twitter_keys.json.example` into a new file called
`config/twitter_keys.json` and set your keys there.
 
Otherwise you can export the keys as environment variables:
- `TWITTER_CONSUMER_KEY`
- `TWITTER_CONSUMER_SECRET`
- `TWITTER_ACCESS_TOKEN`
- `TWITTER_ACCESS_SECRET`

If you're running the app somewhere other than your local machine, you should
only set your keys via environment variables.

## Filter

The Tweets the app will parse are based on the `filter` parameter you can define
in `config/stream_params.json`. It is an array of strings. If tweets are
created matching your provided parameters then they will be pulled in by the
app.

## Running the App

Ensure you have [Docker](https://www.docker.com/) installed on your machine.

To run the app `make run`

To run the tests `make test`


## What is outputted?

Clients (like [Tweetatlas
Client](https://github.com/Ghrehh/tweetatlas-client)) can connect to the server
via a websocket connection. The server will then parse incoming tweets matching
your provided filter parameters and attempt to figure out which country they
originated from. This information is then fed to your client as a JSON object
like the following:

``` json
{
  "IE": 10,
  "US": 100,
  "unknown": 200,
  etc...
}
```

Where the key is the countries [ISO 3166
Code](https://en.wikipedia.org/wiki/ISO_3166) and the value is the number of
tweets associated with that country. Any tweet with a location that can not be
parsed is assigned to the `unknown` key.
