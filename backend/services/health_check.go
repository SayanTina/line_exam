package services

import (
	"net/http"
	"strings"
)

func HealthCheckService(url string, ch chan<- int) {

	// case space in cell or empty return state 3
	web_url := strings.TrimSpace(url)
	if web_url == "" {
		ch <- 3
	}

	// Check url has prefix "http://", Add "http://" to url if missing
	web_url = checkURL(url)

	// call http.Get for check web service
	// return state 0 if cannot reach the website
	_, err := http.Get(web_url)
	if err != nil {
		ch <- 0
	}
	// return state 1 if any http status code that is returned from the website
	ch <- 1
}

// Check url has prefix "http://", Add "http://" to url if missing
func checkURL(url string) string {
	if !strings.HasPrefix(url, "http") {
		return "http://" + url
	}
	return url
}
