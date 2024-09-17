package github

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

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
	res := assert.NoError(http.DefaultClient.Do(newDefaultClient(g.token).get("https://api.github.com/user/starred")))
	body := assert.NoError(io.ReadAll(res.Body))

	repositories := repositories{}
	json.Unmarshal(body, &repositories)

	return repositories
}

func (g GitHub) UserRepos() []Repository {
	res := assert.NoError(http.DefaultClient.Do(newDefaultClient(g.token).get("https://api.github.com/user/repos")))
	body := assert.NoError(io.ReadAll(res.Body))

	repositories := repositories{}
	json.Unmarshal(body, &repositories)

	return repositories
}

func (g GitHub) WatchedRepos() []Repository {
	res := assert.NoError(http.DefaultClient.Do(newDefaultClient(g.token).get("https://api.github.com/user/subscriptions")))
	body := assert.NoError(io.ReadAll(res.Body))

	repositories := repositories{}
	json.Unmarshal(body, &repositories)

	return repositories
}

type defaultClient struct {
	get func(url string) *http.Request
}

func newDefaultClient(token string) defaultClient {
	return defaultClient{
		get: func(url string) *http.Request {
			req := assert.NoError(http.NewRequest("GET", url, nil))

			req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", token))
			req.Header.Add("Accept", "application/vnd.github+json")
			req.Header.Add("X-GitHub-Api-Version", "2022-11-28")

			query := req.URL.Query()
			query.Add("per_page", "100")

			req.URL.RawQuery = query.Encode()

			return req
		},
	}
}
