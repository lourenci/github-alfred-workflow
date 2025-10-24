package repository_test

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
	"github.com/lourenci/github-alfred/usecases/getuseropenpullsinalfred/repository"
	"github.com/stretchr/testify/require"
)

func TestUserOpenPullsOfRepo(t *testing.T) {
	t.Run("returns open pull requests for a given user", func(t *testing.T) {
		fakeHttpClient := test.NewFakeHttpClient(
			map[url.URL][]http.Response{
				*assert.NoError(url.Parse("https://api.github.com/search/issues?q=is:pr+state:open+author:bar")): {
					{
						StatusCode: http.StatusOK,
						Body: io.NopCloser(
							strings.NewReader(
								`{
										"items": [
											{
												"title": "Amazing PR 2",
												"created_at": "2011-01-24T19:01:12Z",
												"html_url": "https://api.github.com/repos/lourenci/foo/pulls/1347",
												"repository_url": "https://api.github.com/repos/lourenci/foo"
											}
										]
									}
						`,
							),
						),
					},
				},
			},
		)

		token := "token"
		pulls := repository.New(github.New(token, fakeHttpClient)).OpenPulls("bar")

		require.Equal(
			t,
			[]github.Pull{
				{
					Title:          "Amazing PR 2",
					CreatedAt:      "2011-01-24T19:01:12Z",
					URL:            "https://api.github.com/repos/lourenci/foo/pulls/1347",
					RepositoryName: "lourenci/foo",
				},
			},
			pulls,
		)
		require.Equal(
			t,
			fakeHttpClient.Calls,
			map[url.URL][]test.Call{
				*assert.NoError(url.Parse("https://api.github.com/search/issues?q=is:pr+state:open+author:bar")): {
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
