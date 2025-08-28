package main

import (
	"fmt"
	"log"
	"net/url"
	"os"
	"strconv"
	"sync"
)

type config struct {
	pages              map[string]int
	baseURL            *url.URL
	mu                 *sync.Mutex
	concurrencyControl chan struct{}
	wg                 *sync.WaitGroup
	maxPages           int
}

func main() {
	if len(os.Args) < 4 {
		log.Fatal("too few arguments provided")
	}
	if len(os.Args) > 4 {
		log.Fatal("too many arguments provided")
	}

	rawBaseUrl := os.Args[1]
	baseUrl, err := url.Parse(rawBaseUrl)
	if err != nil {
		log.Fatalf("couldn't parse base url: %s", rawBaseUrl)
	}

	maxConcurrency, err := strconv.Atoi(os.Args[2])
	if err != nil {
		log.Fatalf("couldn't parse max concurrency choose a number")
	}

	maxPages, err := strconv.Atoi(os.Args[3])
	if err != nil {
		log.Fatalf("couldn't parse max pages choose a number")
	}

	cfg := config{
		pages:              make(map[string]int),
		baseURL:            baseUrl,
		mu:                 &sync.Mutex{},
		concurrencyControl: make(chan struct{}, maxConcurrency),
		wg:                 &sync.WaitGroup{},
		maxPages:           maxPages,
	}

	fmt.Printf("starting crawl of: %s\n", baseUrl)
	cfg.wg.Add(1)
	cfg.crawlPage(rawBaseUrl)
	cfg.wg.Wait()

	fmt.Println(cfg.pages)
}
