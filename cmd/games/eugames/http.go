package main

import (
	"net/http"
	"time"
)

// httpClient is the default http client used for making
// request.
//
// It is just a "doer" so it can be easily replaced with
// clients with retry mechanisms or other policies.
var httpClient doer = newTimeoutClient()

type doer interface {
	Do(*http.Request) (*http.Response, error)
}

type timeoutClient struct {
	*http.Client
}

func newTimeoutClient() timeoutClient {
	return timeoutClient{
		Client: &http.Client{
			Timeout: 10 * time.Second,
		},
	}
}
