package main

import (
	"fmt"
	"log"
	"net/url"
	"os"
	"sync"
)

type config struct {
	pages              map[string]int
	baseURL            *url.URL
	mu                 *sync.Mutex
	concurrencyControl chan struct{}
	wg                 *sync.WaitGroup
}

func main() {
	if len(os.Args) < 2 {
		log.Fatal("no website provided")
	}
	if len(os.Args) > 2 {
		log.Fatal("too many arguments provided")
	}

	rawBaseUrl := os.Args[1]
	baseUrl, err := url.Parse(rawBaseUrl)
	if err != nil {
		log.Fatalf("couldn't parse base url: %s", rawBaseUrl)
	}

	fmt.Printf("starting crawl of: %s\n", baseUrl)

	cfg := config{
		pages:              make(map[string]int),
		baseURL:            baseUrl,
		mu:                 &sync.Mutex{},
		concurrencyControl: make(chan struct{}, 5),
		wg:                 &sync.WaitGroup{},
	}

	cfg.wg.Add(1)
	cfg.crawlPage(rawBaseUrl)
	cfg.wg.Wait()

	fmt.Println(cfg.pages)
}
