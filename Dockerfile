FROM golang:1.9.2-alpine

RUN apk add --no-cache git
RUN go get -u github.com/ghrehh/findlocation
RUN go get -u github.com/dghubble/go-twitter/twitter
RUN go get -u github.com/dghubble/oauth1
RUN go get -u github.com/gorilla/handlers
RUN go get -u github.com/gorilla/mux
RUN go get -u github.com/gorilla/websocket

COPY . ${GOPATH}/src/github.com/ghrehh/tweetatlas

WORKDIR ${GOPATH}/src/github.com/ghrehh/tweetatlas

RUN go install

EXPOSE 8080
CMD tweetatlas
