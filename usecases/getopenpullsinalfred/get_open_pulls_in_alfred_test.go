package getopenpullsinalfred_test

import (
	"testing"

	"github.com/lourenci/github-alfred/lib/github"
	"github.com/lourenci/github-alfred/usecases/getopenpullsinalfred"
	"github.com/lourenci/github-alfred/usecases/getopenpullsinalfred/vo"
	"github.com/stretchr/testify/require"
)

func TestGetUserReposInAlfred(t *testing.T) {
	t.Run("returns all open pr of a repo from a user", func(t *testing.T) {
		{
			alfred := getopenpullsinalfred.New(
				newFakeRepository(
					map[string]map[string][]github.Pull{
						"octocat/Hello-World": {
							"john": []github.Pull{
								{
									Title:     "Amazing PR",
									CreatedAt: "2011-01-26T19:01:12Z",
									URL:       "https://github.com/repos/octocat/Hello-World/pulls/1347",
								},
								{
									Title:     "Amazing PR 2",
									CreatedAt: "2011-01-24T19:01:12Z",
									URL:       "https://github.com/repos/octocat/Hello-World/pulls/1348",
								},
							},
						},
					},
				),
			).GetUserOpenPullsOfRepo("octocat/Hello-World", "john")

			require.Equal(
				t,
				getopenpullsinalfred.Alfred{
					Items: []getopenpullsinalfred.Item{
						{
							UID:      "https://github.com/repos/octocat/Hello-World/pulls/1347",
							Title:    "Amazing PR",
							Subtitle: "Created at 2011-01-26T19:01:12Z",
							Match:    "Amazing PR",
							Arg:      "https://github.com/repos/octocat/Hello-World/pulls/1347",
						},
						{
							UID:      "https://github.com/repos/octocat/Hello-World/pulls/1348",
							Title:    "Amazing PR 2",
							Subtitle: "Created at 2011-01-24T19:01:12Z",
							Match:    "Amazing PR 2",
							Arg:      "https://github.com/repos/octocat/Hello-World/pulls/1348",
						},
					},
				},
				alfred,
			)
		}
		{
			alfred := getopenpullsinalfred.New(
				newFakeRepository(
					map[string]map[string][]github.Pull{
						"foo/bar": {
							"alice": []github.Pull{
								{
									Title:     "Foo PR",
									CreatedAt: "2011-02-26T19:01:12Z",
									URL:       "https://github.com/repos/foo/bar/pulls/1347",
								},
								{
									Title:     "Foo PR 2",
									CreatedAt: "2011-03-24T19:01:12Z",
									URL:       "https://github.com/repos/foo/bar/pulls/1348",
								},
							},
						},
					},
				),
			).GetUserOpenPullsOfRepo("foo/bar", "alice")

			require.Equal(
				t,
				getopenpullsinalfred.Alfred{
					Items: []getopenpullsinalfred.Item{
						{
							UID:      "https://github.com/repos/foo/bar/pulls/1347",
							Title:    "Foo PR",
							Subtitle: "Created at 2011-02-26T19:01:12Z",
							Match:    "Foo PR",
							Arg:      "https://github.com/repos/foo/bar/pulls/1347",
						},
						{
							UID:      "https://github.com/repos/foo/bar/pulls/1348",
							Title:    "Foo PR 2",
							Subtitle: "Created at 2011-03-24T19:01:12Z",
							Match:    "Foo PR 2",
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
	pulls map[string]map[string][]github.Pull
}

func newFakeRepository(pulls map[string]map[string][]github.Pull) fakeRepository {
	return fakeRepository{pulls: pulls}

}

func (f fakeRepository) OpenPulls(repo vo.Repo, user string) []github.Pull {
	return f.pulls[string(repo)][user]
}
