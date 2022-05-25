package services

import (
	"net/http"
	"strings"
)

func HealthCheckService(url string, ch chan<- int) {
	web_url := checkURL(url)

	_, err := http.Get(web_url)
	if err != nil {
		ch <- 0
	}
	ch <- 1
}

func checkURL(url string) string {
	if !strings.HasPrefix(url, "http") {
		return "http://" + url
	}
	return url
}
