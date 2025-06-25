package repository_test

import (
	"io"
	"net/http"
	"net/url"
	"strings"
	"testing"

	"github.com/lourenci/github-alfred/lib/assert"
	"github.com/lourenci/github-alfred/lib/github"
	"github.com/lourenci/github-alfred/lib/http/test"
	"github.com/lourenci/github-alfred/usecases/convertrepositorytoalfred/repository"
	"github.com/stretchr/testify/require"
)

func TestGetAllUserRepos(t *testing.T) {
	t.Run("returns all the starred, watched and user repositories", func(t *testing.T) {
		requests := map[url.URL][]http.Response{
			*assert.NoError(url.Parse("https://api.github.com/user/starred?per_page=100&page=1")): {
				{
					StatusCode: http.StatusOK,
					Body: io.NopCloser(
						strings.NewReader(
							`[
							{
								"full_name": "octocat/starred",
								"html_url": "https://github.com/octocat/starred",
								"description": "starred repo",
								"ssh_url": "git@github.com:octocat/starred.git"
							}
						]
						`,
						),
					),
				},
			},
			*assert.NoError(url.Parse("https://api.github.com/user/repos?per_page=100&page=1")): {
				{StatusCode: http.StatusOK,
					Body: io.NopCloser(
						strings.NewReader(
							`[
							{
								"full_name": "octocat/repos",
								"html_url": "https://github.com/octocat/repos",
								"description": "user repo",
								"ssh_url": "git@github.com:octocat/repos.git"
							}
						]
						`,
						),
					),
				},
			},
			*assert.NoError(url.Parse("https://api.github.com/user/subscriptions?per_page=100&page=1")): {
				{StatusCode: http.StatusOK,
					Body: io.NopCloser(
						strings.NewReader(
							`[
							{
								"full_name": "octocat/subscriptions",
								"html_url": "https://github.com/octocat/subscriptions",
								"description": "subscribed repo",
								"ssh_url": "git@github.com:octocat/subscriptions.git"
							}
						]
						`,
						),
					),
				},
			},
		}

		require.ElementsMatch(
			t,
			[]github.Repository{
				{
					Name:        "octocat/starred",
					URL:         "https://github.com/octocat/starred",
					Description: "starred repo",
					SshURL:      "git@github.com:octocat/starred.git",
				},
				{
					Name:        "octocat/repos",
					URL:         "https://github.com/octocat/repos",
					Description: "user repo",
					SshURL:      "git@github.com:octocat/repos.git",
				},
				{
					Name:        "octocat/subscriptions",
					URL:         "https://github.com/octocat/subscriptions",
					Description: "subscribed repo",
					SshURL:      "git@github.com:octocat/subscriptions.git",
				},
			},
			repository.New(github.New("token", test.NewFakeHttpClient(requests))).GetAllUserRepos(),
		)
	})

	t.Run("ignores duplicated repos", func(t *testing.T) {
		requests := map[url.URL][]http.Response{
			*assert.NoError(url.Parse("https://api.github.com/user/starred?per_page=100&page=1")): {
				{
					StatusCode: http.StatusOK,
					Body: io.NopCloser(
						strings.NewReader(
							`[
							{
								"full_name": "octocat/starred",
								"html_url": "https://github.com/octocat/starred",
								"description": "starred repo",
								"ssh_url": "git@github.com:octocat/starred.git"
							}
						]
						`,
						),
					),
				},
			},
			*assert.NoError(url.Parse("https://api.github.com/user/repos?per_page=100&page=1")): {
				{StatusCode: http.StatusOK,
					Body: io.NopCloser(
						strings.NewReader(
							`[
							{
								"full_name": "octocat/starred",
								"html_url": "https://github.com/octocat/starred",
								"description": "starred repo",
								"ssh_url": "git@github.com:octocat/starred.git"
							}
						]
						`,
						),
					),
				},
			},
			*assert.NoError(url.Parse("https://api.github.com/user/subscriptions?per_page=100&page=1")): {
				{StatusCode: http.StatusOK,
					Body: io.NopCloser(
						strings.NewReader(
							`[
							{
								"full_name": "octocat/starred",
								"html_url": "https://github.com/octocat/starred",
								"description": "starred repo",
								"ssh_url": "git@github.com:octocat/starred.git"
							}
						]
						`,
						),
					),
				},
			},
		}

		require.ElementsMatch(
			t,
			[]github.Repository{
				{
					Name:        "octocat/starred",
					URL:         "https://github.com/octocat/starred",
					Description: "starred repo",
					SshURL:      "git@github.com:octocat/starred.git",
				},
			},
			repository.New(github.New("token", test.NewFakeHttpClient(requests))).GetAllUserRepos(),
		)
	})
}
