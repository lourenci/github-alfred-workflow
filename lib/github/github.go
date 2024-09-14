package github

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type GitHub struct {
	token string
}

type Repositories []Repository

type Repository struct {
	Name        string `json:"full_name"`
	URL         string `json:"html_url"`
	Description string `json:"description"`
	GitURL      string `json:"git_url"`
}

func New(token string) GitHub {
	return GitHub{
		token: token,
	}
}

func (g GitHub) StarredRepos() []Repository {
	res, _ := http.DefaultClient.Do(newDefaultClient(g.token).get("https://api.github.com/user/starred"))
	body, _ := io.ReadAll(res.Body)

	repositories := Repositories{}
	json.Unmarshal(body, &repositories)

	return repositories
}

func (g GitHub) UserRepos() []Repository {
	res, _ := http.DefaultClient.Do(newDefaultClient(g.token).get("https://api.github.com/user/repos"))
	body, _ := io.ReadAll(res.Body)

	repositories := Repositories{}
	json.Unmarshal(body, &repositories)

	return repositories
}

func (g GitHub) WatchedRepos() []Repository {
	res, _ := http.DefaultClient.Do(newDefaultClient(g.token).get("https://api.github.com/user/subscriptions"))
	body, _ := io.ReadAll(res.Body)

	repositories := Repositories{}
	json.Unmarshal(body, &repositories)

	return repositories
}

type defaultClient struct {
	get func(url string) *http.Request
}

func newDefaultClient(token string) defaultClient {
	return defaultClient{
		get: func(url string) *http.Request {
			req, _ := http.NewRequest("GET", url, nil)

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
