package getuseropenpullsinalfred_test

import (
	"testing"

	"github.com/lourenci/github-alfred/lib/github"
	"github.com/lourenci/github-alfred/usecases/getuseropenpullsinalfred"
	"github.com/stretchr/testify/require"
)

func TestGetUserOpenPulls(t *testing.T) {
	t.Run("returns all open pr from a user", func(t *testing.T) {
		{
			alfred := getuseropenpullsinalfred.New(
				newFakeRepository(
					map[string][]github.Pull{
						"john": {
							{
								Title:          "Amazing PR",
								CreatedAt:      "2011-01-26T19:01:12Z",
								URL:            "https://github.com/repos/octocat/Hello-World/pulls/1347",
								RepositoryName: "octocat/Hello-World",
							},
							{
								Title:          "Amazing PR 2",
								CreatedAt:      "2011-01-24T19:01:12Z",
								URL:            "https://github.com/repos/octocat/Hello-World/pulls/1348",
								RepositoryName: "octocat/Hello-World",
							},
						},
					},
				),
			).GetUserOpenPulls("john")

			require.Equal(
				t,
				getuseropenpullsinalfred.Alfred{
					Items: []getuseropenpullsinalfred.Item{
						{
							UID:      "https://github.com/repos/octocat/Hello-World/pulls/1347",
							Title:    "Amazing PR",
							Subtitle: "octocat/Hello-World, created at 2011-01-26T19:01:12Z",
							Match:    "octocat/Hello-World Amazing PR",
							Arg:      "https://github.com/repos/octocat/Hello-World/pulls/1347",
						},
						{
							UID:      "https://github.com/repos/octocat/Hello-World/pulls/1348",
							Title:    "Amazing PR 2",
							Subtitle: "octocat/Hello-World, created at 2011-01-24T19:01:12Z",
							Match:    "octocat/Hello-World Amazing PR 2",
							Arg:      "https://github.com/repos/octocat/Hello-World/pulls/1348",
						},
					},
				},
				alfred,
			)
		}
		{
			alfred := getuseropenpullsinalfred.New(
				newFakeRepository(
					map[string][]github.Pull{
						"alice": {
							{
								Title:          "Foo PR",
								CreatedAt:      "2011-02-26T19:01:12Z",
								URL:            "https://github.com/repos/foo/bar/pulls/1347",
								RepositoryName: "foo/bar",
							},
							{
								Title:          "Foo PR 2",
								CreatedAt:      "2011-03-24T19:01:12Z",
								URL:            "https://github.com/repos/foo/bar/pulls/1348",
								RepositoryName: "foo/bar",
							},
						},
					},
				),
			).GetUserOpenPulls("alice")

			require.Equal(
				t,
				getuseropenpullsinalfred.Alfred{
					Items: []getuseropenpullsinalfred.Item{
						{
							UID:      "https://github.com/repos/foo/bar/pulls/1347",
							Title:    "Foo PR",
							Subtitle: "foo/bar, created at 2011-02-26T19:01:12Z",
							Match:    "foo/bar Foo PR",
							Arg:      "https://github.com/repos/foo/bar/pulls/1347",
						},
						{
							UID:      "https://github.com/repos/foo/bar/pulls/1348",
							Title:    "Foo PR 2",
							Subtitle: "foo/bar, created at 2011-03-24T19:01:12Z",
							Match:    "foo/bar Foo PR 2",
							Arg:      "https://github.com/repos/foo/bar/pulls/1348",
						},
					},
				},
				alfred,
			)
		}
	})
}

type fakeRepository struct {
	pulls map[string][]github.Pull
}

func newFakeRepository(pulls map[string][]github.Pull) fakeRepository {
	return fakeRepository{pulls: pulls}

}

func (f fakeRepository) OpenPulls(user string) []github.Pull {
	return f.pulls[user]
}
