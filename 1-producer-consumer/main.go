//////////////////////////////////////////////////////////////////////
//
// Given is a producer-consumer szenario, where a producer reads in
// tweets from a mockstream and a consumer is processing the
// data. Your task is to change the code so that the producer as well
// as the consumer can run concurrently
//

package main

import (
	"fmt"
	"time"

	"github.com/kennykarnama/go-concurrency-exercises/1-producer-consumer/procon"
)

func producer(stream procon.Stream) (tweets []*procon.Tweet) {
	for {
		tweet, err := stream.Next()
		if err == procon.ErrEOF {
			return tweets
		}

		tweets = append(tweets, tweet)
	}
}

func consumer(tweets []*procon.Tweet) {
	for _, t := range tweets {
		if t.IsTalkingAboutGo() {
			fmt.Println(t.Username, "\ttweets about golang")
		} else {
			fmt.Println(t.Username, "\tdoes not tweet about golang")
		}
	}
}

func main() {
	start := time.Now()
	stream := procon.GetMockStream()

	// Producer
	tweets := producer(stream)

	// Consumer
	consumer(tweets)

	fmt.Printf("Process took %s\n", time.Since(start))
}
