package main

import (
	"encoding/json"
	"flag"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"

	"code.google.com/p/go.net/websocket"
	"github.com/darkhelmet/twitterstream"
)

type Tweet struct {
	Hashtags   []string `json:"hashtags"`
	Mentions   []string `json:"mentions"`
	Text       string   `json:"text"`
	Name       string   `json:"name"`
	ScreenName string   `json:"screenName"`
}

type Client struct {
	conn      *websocket.Conn
	closeConn chan bool
}

var (
	keywords       = flag.String("keywords", "", "keywords to track")
	wait           = 1
	maxWait        = 600 // Seconds
	consumerKey    string
	consumerSecret string
	accessToken    string
	accessSecret   string
	tweets         = make(chan *Tweet, 10)
	clients        = make(chan *Client)
)

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func init() {
	consumerKey = os.Getenv("TWITTER_CONSUMER_KEY")
	consumerSecret = os.Getenv("TWITTER_CONSUMER_SECRET")
	accessToken = os.Getenv("TWITTER_ACCESS_TOKEN")
	accessSecret = os.Getenv("TWITTER_ACCESS_SECRET")

	if consumerKey == "" || consumerSecret == "" {
		log.Fatalln("consumer tokens left blank")
	}

	if accessToken == "" || accessSecret == "" {
		log.Fatalln("access tokens left blank")
	}

	flag.Parse()
	if *keywords == "" {
		log.Fatalln("keywords left blank")
	}

	go wsConnHandler()
	go downloader()
}

func decodeTweets(conn *twitterstream.Connection) {
	var tweet *Tweet
	for {
		if retrievedTweet, err := conn.Next(); err == nil {
			log.Printf("New Tweet: %s\n", retrievedTweet.Text)
			tweet = new(Tweet)
			for _, hashtag := range retrievedTweet.Entities.Hashtags {
				tweet.Hashtags = append(tweet.Hashtags, hashtag.Text)
			}
			for _, mention := range retrievedTweet.Entities.Mentions {
				tweet.Mentions = append(tweet.Mentions, mention.ScreenName)
			}
			tweet.Text = retrievedTweet.Text
			tweet.Name = retrievedTweet.User.Name
			tweet.ScreenName = retrievedTweet.User.ScreenName

			tweets <- tweet
		} else {
			log.Printf("decoding tweet failed: %s", err)
			return
		}
	}
}

func downloader() {
	client := twitterstream.NewClient(consumerKey, consumerSecret, accessToken, accessSecret)
	for {
		log.Printf("tracking keywords %s", *keywords)
		conn, err := client.Track(*keywords)
		defer conn.Close()
		if err != nil {
			log.Printf("tracking failed: %s", err)
			wait = wait << 1
			log.Printf("waiting for %d seconds before reconnect", min(wait, maxWait))
			time.Sleep(time.Duration(min(wait, maxWait)) * time.Second)
			continue
		} else {
			wait = 1
		}
		decodeTweets(conn)

	}
}

func wsConnHandler() {
	conns := make(map[*websocket.Conn]*Client)
	for {
		select {
		case client := <-clients:
			conns[client.conn] = client
		case tweet := <-tweets:
			for conn, client := range conns {
				json_tweet, _ := json.Marshal(tweet)
				if _, err := conn.Write(json_tweet); err != nil {
					conn.Close()
					client.closeConn <- true
					close(client.closeConn)
				}
			}
		}
	}
}

func IndexHandler(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "text/html")
	file, err := ioutil.ReadFile("index.html")
	if err != nil {
		panic(err)
	}
	res.Write(file)
}

func TweetsHandler(ws *websocket.Conn) {
	log.Println("Client Connected")
	client := &Client{ws, make(chan bool)}
	clients <- client
	<-client.closeConn
}

func main() {
	server := http.NewServeMux()

	server.HandleFunc("/", IndexHandler)
	server.Handle("/ws", websocket.Handler(TweetsHandler))

	err := http.ListenAndServe(":3000", server)
	if err != nil {
		panic(err)
	}
}
