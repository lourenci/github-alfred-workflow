package test

import (
	"net/http"
	"net/url"
	"sync"
)

type Call struct {
	Headers map[string]string
}

type fakeHttpClient struct {
	lock     sync.Mutex
	Calls    map[url.URL][]Call
	requests map[url.URL][]http.Response
}

func NewFakeHttpClient(requests map[url.URL][]http.Response) *fakeHttpClient {
	return &fakeHttpClient{
		Calls:    make(map[url.URL][]Call),
		requests: requests,
	}
}

func (f *fakeHttpClient) Get(url url.URL, headers map[string]string) (*http.Response, error) {
	f.lock.Lock()
	defer f.lock.Unlock()
	f.Calls[url] = append(f.Calls[url], Call{Headers: headers})

	return &f.requests[url][len(f.Calls[url])-1], nil
}
