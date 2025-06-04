package github

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"regexp"
	"strconv"

	"github.com/lourenci/github-alfred/lib/assert"
)

type GitHub struct {
	token string
}

type repositories []Repository

type Repository struct {
	Name        string `json:"full_name"`
	URL         string `json:"html_url"`
	Description string `json:"description"`
	SshURL      string `json:"ssh_url"`
}

func New(token string) GitHub {
	return GitHub{
		token: token,
	}
}

func (g GitHub) StarredRepos() []Repository {
	all_repositories := repositories{}

	responses := assert.NoError(newDefaultClient(g.token).get("https://api.github.com/user/starred"))
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

	responses := assert.NoError(newDefaultClient(g.token).get("https://api.github.com/user/repos"))
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

	responses := assert.NoError(newDefaultClient(g.token).get("https://api.github.com/user/subscriptions"))
	for _, res := range responses {
		body := assert.NoError(io.ReadAll(res.Body))

		page_repositories := repositories{}
		json.Unmarshal(body, &page_repositories)

		all_repositories = append(all_repositories, page_repositories...)
	}

	return all_repositories
}

type defaultClient struct {
	token string
}

func newDefaultClient(token string) defaultClient {
	return defaultClient{
		token: token,
	}
}

func (c defaultClient) get(url string) ([]*http.Response, error) {
	return c.get_with_page(url, 1, []*http.Response{})
}

func (c defaultClient) get_with_page(url string, page int, responses []*http.Response) ([]*http.Response, error) {
	req := assert.NoError(http.NewRequest("GET", url, nil))

	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", c.token))
	req.Header.Add("Accept", "application/vnd.github+json")
	req.Header.Add("X-GitHub-Api-Version", "2022-11-28")

	query := req.URL.Query()
	query.Add("per_page", "100")
	query.Add("page", fmt.Sprintf("%d", page))

	req.URL.RawQuery = query.Encode()

	res := assert.NoError(http.DefaultClient.Do(req))

	responses = append(responses, res)

	lastPage := page
	if matches := regexp.MustCompile(`rel="next".*\bpage\b=(\d+).* rel="last"`).FindStringSubmatch(res.Header.Get("link")); len(matches) > 0 {
		lastPage = assert.NoError(strconv.Atoi(matches[1]))
	}

	if lastPage == page {
		return responses, nil
	}

	return c.get_with_page(url, page+1, responses)
}
