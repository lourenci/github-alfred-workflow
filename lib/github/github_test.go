package github_test

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
	"testing"

	"github.com/lourenci/github-alfred/lib/assert"
	"github.com/lourenci/github-alfred/lib/github"
	"github.com/lourenci/github-alfred/lib/http/test"
	"github.com/stretchr/testify/require"
)

func TestStarredRepos(t *testing.T) {
	t.Run("returns a list of starred repositories", func(t *testing.T) {
		fakeHttpClient := test.NewFakeHttpClient(
			map[url.URL][]http.Response{
				*assert.NoError(url.Parse("https://api.github.com/user/starred?per_page=100&page=1")): {
					{
						StatusCode: http.StatusOK,
						Body: io.NopCloser(
							strings.NewReader(
								`[
							{
								"full_name": "octocat/Hello-World",
								"html_url": "https://github.com/octocat/Hello-World",
								"description": "This your first repo!",
								"ssh_url": "git@github.com:octocat/Hello-World.git"
							}
						]
						`,
							),
						),
					},
				},
			},
		)
		token := "test-token"

		githubClient := github.New(token, fakeHttpClient)

		repo := githubClient.StarredRepos()

		require.Equal(
			t,
			[]github.Repository{
				{
					Name:        "octocat/Hello-World",
					URL:         "https://github.com/octocat/Hello-World",
					Description: "This your first repo!",
					SshURL:      "git@github.com:octocat/Hello-World.git",
				},
			},
			repo,
		)
		require.Equal(
			t,
			fakeHttpClient.Calls,
			map[url.URL][]test.Call{
				*assert.NoError(url.Parse("https://api.github.com/user/starred?per_page=100&page=1")): {
					{
						Headers: map[string]string{
							"Authorization":        fmt.Sprintf("Bearer %s", token),
							"Accept":               "application/vnd.github+json",
							"X-GitHub-Api-Version": "2022-11-28",
						},
					},
				},
			},
		)
	})

	t.Run("returns all pages of starred repositories", func(t *testing.T) {
		headerResponse1 := http.Header{}
		headerResponse1.Set("link", `<https://api.github.com/user/starred?page=2>; rel="next", <https://api.github.com/user/starred?page=2>; rel="last", <https://api.github.com/user/starred?page=1>; rel="first"`)
		headerResponse2 := http.Header{}
		headerResponse2.Set("link", `<https://api.github.com/user/starred?page=1>; rel="prev", <https://api.github.com/user/starred?page=2>; rel="last", <https://api.github.com/user/starred?page=1>; rel="first"`)
		fakeHttpClient := test.NewFakeHttpClient(
			map[url.URL][]http.Response{
				*assert.NoError(url.Parse("https://api.github.com/user/starred?per_page=100&page=1")): {
					{
						StatusCode: http.StatusOK,
						Header:     headerResponse1,
						Body: io.NopCloser(
							strings.NewReader(
								`[
							{
								"full_name": "octocat/Hello-World",
								"html_url": "https://github.com/octocat/Hello-World",
								"description": "This your first repo!",
								"ssh_url": "git@github.com:octocat/Hello-World.git"
							}
						]
						`,
							),
						),
					},
				},
				*assert.NoError(url.Parse("https://api.github.com/user/starred?per_page=100&page=2")): {
					{
						StatusCode: http.StatusOK,
						Header:     headerResponse2,
						Body: io.NopCloser(
							strings.NewReader(
								`[
							{
								"full_name": "lourenci/alfred-github",
								"html_url": "https://github.com/lourenci/alfred-github",
								"description": "Alfred github workflow",
								"ssh_url": "git@github.com:lourenci/alfred-github.git"
							}
						]
						`,
							),
						),
					},
				},
			})
		token := "test-token"

		githubClient := github.New(token, fakeHttpClient)

		repo := githubClient.StarredRepos()

		require.Equal(
			t,
			[]github.Repository{
				{
					Name:        "octocat/Hello-World",
					URL:         "https://github.com/octocat/Hello-World",
					Description: "This your first repo!",
					SshURL:      "git@github.com:octocat/Hello-World.git",
				},
				{
					Name:        "lourenci/alfred-github",
					URL:         "https://github.com/lourenci/alfred-github",
					Description: "Alfred github workflow",
					SshURL:      "git@github.com:lourenci/alfred-github.git",
				},
			},
			repo,
		)
		require.Equal(
			t,
			fakeHttpClient.Calls,
			map[url.URL][]test.Call{
				*assert.NoError(url.Parse("https://api.github.com/user/starred?per_page=100&page=1")): {
					{
						Headers: map[string]string{
							"Authorization":        fmt.Sprintf("Bearer %s", token),
							"Accept":               "application/vnd.github+json",
							"X-GitHub-Api-Version": "2022-11-28",
						},
					},
				},
				*assert.NoError(url.Parse("https://api.github.com/user/starred?per_page=100&page=2")): {
					{
						Headers: map[string]string{
							"Authorization":        fmt.Sprintf("Bearer %s", token),
							"Accept":               "application/vnd.github+json",
							"X-GitHub-Api-Version": "2022-11-28",
						},
					},
				},
			},
		)
	})
}

func TestUserRepos(t *testing.T) {
	t.Run("returns a list of user repositories", func(t *testing.T) {
		fakeHttpClient := test.NewFakeHttpClient(
			map[url.URL][]http.Response{
				*assert.NoError(url.Parse("https://api.github.com/user/repos?per_page=100&page=1")): {
					{StatusCode: http.StatusOK,
						Body: io.NopCloser(
							strings.NewReader(
								`[
							{
								"full_name": "octocat/Hello-World",
								"html_url": "https://github.com/octocat/Hello-World",
								"description": "This your first repo!",
								"ssh_url": "git@github.com:octocat/Hello-World.git"
							}
						]
						`,
							),
						),
					},
				},
			},
		)
		token := "test-token"

		githubClient := github.New(token, fakeHttpClient)

		repo := githubClient.UserRepos()

		require.Equal(
			t,
			[]github.Repository{
				{
					Name:        "octocat/Hello-World",
					URL:         "https://github.com/octocat/Hello-World",
					Description: "This your first repo!",
					SshURL:      "git@github.com:octocat/Hello-World.git",
				},
			},
			repo,
		)
		require.Equal(
			t,
			fakeHttpClient.Calls,
			map[url.URL][]test.Call{
				*assert.NoError(url.Parse("https://api.github.com/user/repos?per_page=100&page=1")): {
					{
						Headers: map[string]string{
							"Authorization":        fmt.Sprintf("Bearer %s", token),
							"Accept":               "application/vnd.github+json",
							"X-GitHub-Api-Version": "2022-11-28",
						},
					},
				},
			},
		)
	})

	t.Run("returns all pages of user repositories", func(t *testing.T) {
		headerResponse1 := http.Header{}
		headerResponse1.Set("link", `<https://api.github.com/user/repos?page=2>; rel="next", <https://api.github.com/user/repos?page=2>; rel="last", <https://api.github.com/user/repos?page=1>; rel="first"`)

		headerResponse2 := http.Header{}
		headerResponse2.Set("link", `<https://api.github.com/user/repos?page=1>; rel="prev", <https://api.github.com/user/repos?page=2>; rel="last", <https://api.github.com/user/repos?page=1>; rel="first"`)

		fakeHttpClient := test.NewFakeHttpClient(
			map[url.URL][]http.Response{
				*assert.NoError(url.Parse("https://api.github.com/user/repos?per_page=100&page=1")): {
					{
						StatusCode: http.StatusOK,
						Header:     headerResponse1,
						Body: io.NopCloser(
							strings.NewReader(
								`[
							{
								"full_name": "octocat/Hello-World",
								"html_url": "https://github.com/octocat/Hello-World",
								"description": "This your first repo!",
								"ssh_url": "git@github.com:octocat/Hello-World.git"
							}
						]
						`,
							),
						),
					},
				},
				*assert.NoError(url.Parse("https://api.github.com/user/repos?per_page=100&page=2")): {
					{
						StatusCode: http.StatusOK,
						Header:     headerResponse2,
						Body: io.NopCloser(
							strings.NewReader(
								`[
							{
								"full_name": "lourenci/alfred-github",
								"html_url": "https://github.com/lourenci/alfred-github",
								"description": "Alfred github workflow",
								"ssh_url": "git@github.com:lourenci/alfred-github.git"
							}
						]
						`,
							),
						),
					},
				},
			},
		)
		token := "test-token"

		githubClient := github.New(token, fakeHttpClient)

		repo := githubClient.UserRepos()

		require.Equal(
			t,
			[]github.Repository{
				{
					Name:        "octocat/Hello-World",
					URL:         "https://github.com/octocat/Hello-World",
					Description: "This your first repo!",
					SshURL:      "git@github.com:octocat/Hello-World.git",
				},
				{
					Name:        "lourenci/alfred-github",
					URL:         "https://github.com/lourenci/alfred-github",
					Description: "Alfred github workflow",
					SshURL:      "git@github.com:lourenci/alfred-github.git",
				},
			},
			repo,
		)
		require.Equal(
			t,
			fakeHttpClient.Calls,
			map[url.URL][]test.Call{
				*assert.NoError(url.Parse("https://api.github.com/user/repos?per_page=100&page=1")): {
					{
						Headers: map[string]string{
							"Authorization":        fmt.Sprintf("Bearer %s", token),
							"Accept":               "application/vnd.github+json",
							"X-GitHub-Api-Version": "2022-11-28",
						},
					},
				},
				*assert.NoError(url.Parse("https://api.github.com/user/repos?per_page=100&page=2")): {
					{
						Headers: map[string]string{
							"Authorization":        fmt.Sprintf("Bearer %s", token),
							"Accept":               "application/vnd.github+json",
							"X-GitHub-Api-Version": "2022-11-28",
						},
					},
				},
			},
		)
	})
}

func TestWatchedRepos(t *testing.T) {
	t.Run("returns a list of user watched repositories", func(t *testing.T) {
		fakeHttpClient := test.NewFakeHttpClient(
			map[url.URL][]http.Response{
				*assert.NoError(url.Parse("https://api.github.com/user/subscriptions?per_page=100&page=1")): {
					{StatusCode: http.StatusOK,
						Body: io.NopCloser(
							strings.NewReader(
								`[
							{
								"full_name": "octocat/Hello-World",
								"html_url": "https://github.com/octocat/Hello-World",
								"description": "This your first repo!",
								"ssh_url": "git@github.com:octocat/Hello-World.git"
							}
						]
						`,
							),
						),
					},
				},
			},
		)
		token := "test-token"

		githubClient := github.New(token, fakeHttpClient)

		repo := githubClient.WatchedRepos()

		require.Equal(
			t,
			[]github.Repository{
				{
					Name:        "octocat/Hello-World",
					URL:         "https://github.com/octocat/Hello-World",
					Description: "This your first repo!",
					SshURL:      "git@github.com:octocat/Hello-World.git",
				},
			},
			repo,
		)
		require.Equal(
			t,
			fakeHttpClient.Calls,
			map[url.URL][]test.Call{
				*assert.NoError(url.Parse("https://api.github.com/user/subscriptions?per_page=100&page=1")): {
					{
						Headers: map[string]string{
							"Authorization":        fmt.Sprintf("Bearer %s", token),
							"Accept":               "application/vnd.github+json",
							"X-GitHub-Api-Version": "2022-11-28",
						},
					},
				},
			},
		)
	})

	t.Run("returns all pages of user watched repositories", func(t *testing.T) {
		headerResponse1 := http.Header{}
		headerResponse1.Set("link", `<https://api.github.com/user/subscriptions?page=2>; rel="next", <https://api.github.com/user/subscriptions?page=2>; rel="last", <https://api.github.com/user/subscriptions?page=1>; rel="first"`)

		headerResponse2 := http.Header{}
		headerResponse2.Set("link", `<https://api.github.com/user/subscriptions?page=1>; rel="prev", <https://api.github.com/user/subscriptions?page=2>; rel="last", <https://api.github.com/user/subscriptions?page=1>; rel="first"`)

		fakeHttpClient := test.NewFakeHttpClient(
			map[url.URL][]http.Response{
				*assert.NoError(url.Parse("https://api.github.com/user/subscriptions?per_page=100&page=1")): {
					{
						StatusCode: http.StatusOK,
						Header:     headerResponse1,
						Body: io.NopCloser(
							strings.NewReader(
								`[
									{
										"full_name": "octocat/Hello-World",
										"html_url": "https://github.com/octocat/Hello-World",
										"description": "This your first repo!",
										"ssh_url": "git@github.com:octocat/Hello-World.git"
									}
								]
								`,
							),
						),
					},
				},
				*assert.NoError(url.Parse("https://api.github.com/user/subscriptions?per_page=100&page=2")): {
					{
						StatusCode: http.StatusOK,
						Header:     headerResponse2,
						Body: io.NopCloser(
							strings.NewReader(
								`[
									{
										"full_name": "lourenci/alfred-github",
										"html_url": "https://github.com/lourenci/alfred-github",
										"description": "Alfred github workflow",
										"ssh_url": "git@github.com:lourenci/alfred-github.git"
									}
								]
								`,
							),
						),
					},
				},
			},
		)
		token := "test-token"

		githubClient := github.New(token, fakeHttpClient)

		repo := githubClient.WatchedRepos()

		require.Equal(
			t,
			[]github.Repository{
				{
					Name:        "octocat/Hello-World",
					URL:         "https://github.com/octocat/Hello-World",
					Description: "This your first repo!",
					SshURL:      "git@github.com:octocat/Hello-World.git",
				},
				{
					Name:        "lourenci/alfred-github",
					URL:         "https://github.com/lourenci/alfred-github",
					Description: "Alfred github workflow",
					SshURL:      "git@github.com:lourenci/alfred-github.git",
				},
			},
			repo,
		)
		require.Equal(
			t,
			fakeHttpClient.Calls,
			map[url.URL][]test.Call{
				*assert.NoError(url.Parse("https://api.github.com/user/subscriptions?per_page=100&page=1")): {
					{
						Headers: map[string]string{
							"Authorization":        fmt.Sprintf("Bearer %s", token),
							"Accept":               "application/vnd.github+json",
							"X-GitHub-Api-Version": "2022-11-28",
						},
					},
				},
				*assert.NoError(url.Parse("https://api.github.com/user/subscriptions?per_page=100&page=2")): {
					{
						Headers: map[string]string{
							"Authorization":        fmt.Sprintf("Bearer %s", token),
							"Accept":               "application/vnd.github+json",
							"X-GitHub-Api-Version": "2022-11-28",
						},
					},
				},
			},
		)
	})
}
