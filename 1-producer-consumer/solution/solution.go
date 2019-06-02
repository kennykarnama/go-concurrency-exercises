package main

import (
	"fmt"
	"sync"
	"time"

	"github.com/kennykarnama/go-concurrency-exercises/1-producer-consumer/procon"
)

//Producer will fill the channel with tweets
func Producer(stream procon.Stream, in chan *procon.Tweet) {
	// Don't forget to close the channel
	defer close(in)
	for {
		tweet, err := stream.Next()

		if err == procon.ErrEOF {
			break
		}
		//Thread block
		in <- tweet
	}
}

//Consumer will consume the tweets given by producer
func Consumer(tweets chan *procon.Tweet, wg *sync.WaitGroup) {

	defer wg.Done()

	for t := range tweets {
		if t.IsTalkingAboutGo() {
			fmt.Println(t.Username, "\ttweets about golang")
		} else {
			fmt.Println(t.Username, "\tdoes not tweet about golang")
		}
	}
}

func main() {
	start := time.Now()
	//Get tweets
	stream := procon.GetMockStream()
	//Construct tweets channel
	in := make(chan *procon.Tweet)
	//Construct waitgroup sync
	wg := &sync.WaitGroup{}

	go Producer(stream, in)
	//How many go routines
	numProcess := 10

	for i := 0; i < numProcess; i++ {
		wg.Add(1)
		go Consumer(in, wg)
	}

	//Wait for all goroutines to finish
	wg.Wait()

	fmt.Printf("Process took %s\n", time.Since(start))
}
