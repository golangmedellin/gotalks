TwitterStreamingExample
=======================

Golang example for meetup with using the twitter streaming API and websockets

set your own environment
```bash
export TWITTER_CONSUMER_KEY="YOUR_CONSUMER_KEY"
export TWITTER_CONSUMER_SECRET="YOUR_CONSUMER_SECRET"
export TWITTER_ACCESS_TOKEN="YOUR_ACCESS_TOKEN"
export TWITTER_ACCESS_SECRET="YOUR_ACCESS_SECRET"
```
then

```bash
# compile golang source code and build the binary executable named go-med
go build -o go-med -race golang_medellin.go
# then exec it!
./go-med -keywords="medellin"
```

