package main

import (
	"fmt"
	"net/url"
	"strings"

	"golang.org/x/net/html"
	"golang.org/x/net/html/atom"
)

func getURLsFromHTML(htmlBody, rawBaseURL string) ([]string, error) {
	baseURL, err := url.Parse(rawBaseURL)
	if err != nil {
		return nil, fmt.Errorf("Couldn't parse base URL")
	}

	doc, err := html.Parse(strings.NewReader(htmlBody))
	if err != nil {
		return nil, fmt.Errorf("Couldn't parse HTML")
	}

	urls := []string{}
	for node := range doc.Descendants() {
		if node.Type == html.ElementNode && node.DataAtom == atom.A {
			for _, a := range node.Attr {
				if a.Key == "href" {
					if a.Val == "" {
						break
					}

					fullUrl, err := joinUrl(*baseURL, a.Val)
					if err != nil {
						return nil, err
					}

					urls = append(urls, fullUrl)
					break
				}
			}
		}
	}

	return urls, nil
}

func joinUrl(baseUrl url.URL, urlPath string) (string, error) {
	u, err := url.Parse(urlPath)
	if err != nil {
		return "", err
	}

	if u.Hostname() != "" {
		return urlPath, nil
	}

	return baseUrl.JoinPath(urlPath).String(), nil
}
