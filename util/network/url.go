package network

import (
	"api-authenticator-proxy/util/log"
	"net/http"
	"net/url"
	"time"
)

func IsURLValid(inputURL string) bool {
	// Parse the URL to check if it is in a valid format
	parsedURL, err := url.ParseRequestURI(inputURL)
	if err != nil {
		log.Warning("invalid URL format: ", inputURL)
		return false
	}

	// Send a HEAD request to check if the URL is reachable
	client := &http.Client{
		Timeout: time.Second * 5, // Set a timeout for the request
	}
	resp, err := client.Head(parsedURL.String())
	if err != nil {
		log.Debug(err)
		log.Warning("the url ", inputURL, " is not reachable")
		return false
	}
	defer resp.Body.Close()

	// Check if the status code indicates success
	if resp.StatusCode >= 200 && resp.StatusCode < 400 {
		return true
	}

	return false
}
