package main

import (
	"fmt"
	"io"
	"net/http"
	"strings"
)

func getHTML(rawURL string) (string, error) {
	req, err := http.NewRequest("GET", rawURL, nil)
	if err != nil {
		return "", fmt.Errorf("Error creating request for: %s", rawURL)
	}

	req.Header.Set("User-Agent", "qwerty")

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return "", fmt.Errorf("Error getting html from: %s. %w", rawURL, err)
	}
	defer res.Body.Close()

	if res.StatusCode >= 400 {
		return "", fmt.Errorf("Error status code: %d", res.StatusCode)
	}

	contentType := res.Header.Get("Content-type")
	if !strings.Contains(contentType, "text/html") {
		return "", fmt.Errorf("Wrong Content-type header: %s", contentType)
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return "", fmt.Errorf("Failed to read response body")
	}

	return string(body), nil
}
