package http

import (
	"net/http"
	"net/url"
)

type httpClient struct{}

func New() httpClient {
	return httpClient{}
}

func (h httpClient) Get(url url.URL, headers map[string]string) (*http.Response, error) {
	client := http.Client{}

	req, _ := http.NewRequest("GET", url.String(), nil)

	for key, value := range headers {
		req.Header.Add(key, value)
	}

	return client.Do(req)
}
