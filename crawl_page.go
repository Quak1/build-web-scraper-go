package main

import (
	"fmt"
	"log"
	"net/url"
)

func (cfg *config) crawlPage(rawCurrentURL string) {
	cfg.concurrencyControl <- struct{}{}
	defer func() {
		<-cfg.concurrencyControl
		cfg.wg.Done()
	}()

	cfg.mu.Lock()
	if len(cfg.pages) >= cfg.maxPages {
		cfg.mu.Unlock()
		return
	}
	cfg.mu.Unlock()

	currentURL, err := url.Parse(rawCurrentURL)
	if err != nil {
		log.Fatalf("Couldn't parse url: %s", rawCurrentURL)
	}

	if cfg.baseURL.Hostname() != currentURL.Hostname() {
		return
	}

	normalizedURL, err := normalizeURL(rawCurrentURL)
	if err != nil {
		log.Fatalf("couldn't normalize url: %s", rawCurrentURL)
	}

	if isFirst := cfg.addPageVisit(normalizedURL); !isFirst {
		return
	}

	htmlBody, err := getHTML(rawCurrentURL)
	if err != nil {
		log.Printf("error getting html for: %s. %v", rawCurrentURL, err)
	}

	urls, err := getURLsFromHTML(htmlBody, cfg.baseURL.String())
	if err != nil {
		log.Fatal(err)
	}

	for _, url := range urls {
		fmt.Println(url)
		cfg.wg.Add(1)
		go cfg.crawlPage(url)
	}
}

func (cfg *config) addPageVisit(normalizedURL string) (isFirst bool) {
	cfg.mu.Lock()
	defer cfg.mu.Unlock()

	if _, visited := cfg.pages[normalizedURL]; visited {
		cfg.pages[normalizedURL]++
		return false
	}

	cfg.pages[normalizedURL] += 1
	return true
}
