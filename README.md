# Tweetatlas Server

## Twitter Keys

To run this app, you must have a set of Twitter API keys. Please refer to the
[Twitter docs](https://developer.twitter.com/en/docs) for information on how to
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
