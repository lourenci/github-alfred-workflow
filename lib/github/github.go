package github

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"regexp"
	"strconv"

	"github.com/lourenci/github-alfred/lib/assert"
)

type GitHub struct {
	token      string
	httpClient HttpClient
}

type repositories []Repository

type Repository struct {
	Name        string `json:"full_name"`
	URL         string `json:"html_url"`
	Description string `json:"description"`
	SshURL      string `json:"ssh_url"`
}

type HttpClient interface {
	Get(url url.URL, headers map[string]string) (*http.Response, error)
}

func New(token string, httpClient HttpClient) GitHub {
	return GitHub{
		token:      token,
		httpClient: httpClient,
	}
}

func (g GitHub) StarredRepos() []Repository {
	all_repositories := repositories{}

	responses := assert.NoError(newDefaultClient(g.httpClient, g.token).get("https://api.github.com/user/starred"))
	for _, res := range responses {
		body := assert.NoError(io.ReadAll(res.Body))

		page_repositories := repositories{}
		json.Unmarshal(body, &page_repositories)

		all_repositories = append(all_repositories, page_repositories...)
	}

	return all_repositories
}

func (g GitHub) UserRepos() []Repository {
	all_repositories := repositories{}

	responses := assert.NoError(newDefaultClient(g.httpClient, g.token).get("https://api.github.com/user/repos"))
	for _, res := range responses {
		body := assert.NoError(io.ReadAll(res.Body))

		page_repositories := repositories{}
		json.Unmarshal(body, &page_repositories)

		all_repositories = append(all_repositories, page_repositories...)
	}

	return all_repositories
}

func (g GitHub) WatchedRepos() []Repository {
	all_repositories := repositories{}

	responses := assert.NoError(newDefaultClient(g.httpClient, g.token).get("https://api.github.com/user/subscriptions"))
	for _, res := range responses {
		body := assert.NoError(io.ReadAll(res.Body))

		page_repositories := repositories{}
		json.Unmarshal(body, &page_repositories)

		all_repositories = append(all_repositories, page_repositories...)
	}

	return all_repositories
}

type defaultClient struct {
	httpClient HttpClient
	token      string
}

func newDefaultClient(httpClient HttpClient, token string) defaultClient {
	return defaultClient{
		httpClient: httpClient,
		token:      token,
	}
}

func (c defaultClient) get(url string) ([]*http.Response, error) {
	return c.get_with_page(url, 1, []*http.Response{})
}

func (c defaultClient) get_with_page(urlString string, page int, responses []*http.Response) ([]*http.Response, error) {
	res := assert.NoError(c.httpClient.Get(
		*assert.NoError(url.Parse(fmt.Sprintf("%s?per_page=%d&page=%d", urlString, 100, page))),
		map[string]string{
			"Authorization":        fmt.Sprintf("Bearer %s", c.token),
			"Accept":               "application/vnd.github+json",
			"X-GitHub-Api-Version": "2022-11-28",
		},
	))

	responses = append(responses, res)

	lastPage := page
	if matches := regexp.MustCompile(`rel="next".*\bpage\b=(\d+).* rel="last"`).FindStringSubmatch(res.Header.Get("link")); len(matches) > 0 {
		lastPage = assert.NoError(strconv.Atoi(matches[1]))
	}

	if lastPage == page {
		return responses, nil
	}

	return c.get_with_page(urlString, page+1, responses)
}
