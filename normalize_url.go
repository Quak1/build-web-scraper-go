package main

import (
	"fmt"
	"net/url"
	"strings"
)

func normalizeURL(urlString string) (string, error) {
	u, err := url.Parse(urlString)

	if err != nil || u.Hostname() == "" {
		return "", fmt.Errorf("invalid url")
	}

	normalized := u.Hostname() + u.Path
	normalized = strings.TrimSuffix(normalized, "/")
	normalized = strings.ToLower(normalized)

	return normalized, nil
}
