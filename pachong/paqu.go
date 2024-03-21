package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
	"sync"
	"time"
)

// Maximum number of working goroutines
const MaxWorkNum = 10

// URL channel
var UrlChannel = make(chan string, MaxWorkNum)

// Results channel
var ResultsChannel = make(chan string, MaxWorkNum)

func GenerateUrlProducer(seedUrl string) {
	go func() {
		UrlChannel <- seedUrl
	}()
}

func GenerateWorkers() {
	var wg sync.WaitGroup
	// Limit the number of working goroutines
	for i := 0; i < MaxWorkNum; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			for {
				url, ok := <-UrlChannel
				if !ok {
					return
				}
				newUrls, err := fetch(url)
				if err != nil {
					fmt.Printf("Worker %d: %v\n", i, err)
					return
				}
				for _, newUrl := range newUrls {
					UrlChannel <- newUrl
				}
				ResultsChannel <- url
			}
		}(i)
	}
	wg.Wait()
}

func fetch(url string) ([]string, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	body, _ := ioutil.ReadAll(resp.Body)
	urlRegexp := regexp.MustCompile(`http[s]?://[^"'\s]+`)
	newUrls := urlRegexp.FindAllString(string(body), -1)
	resp.Body.Close()
	time.Sleep(time.Second) // to prevent IP being blocked
	return newUrls, nil
}

func ResultsConsumer() {
	for {
		url, ok := <-ResultsChannel
		if !ok {
			return
		}
		fmt.Println("Fetched:", url)
	}
}

func main() {
	go GenerateUrlProducer("http://example.com")
	go GenerateWorkers()
	ResultsConsumer()
}
