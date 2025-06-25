package convertrepositorytoalfred_test

import (
	"testing"
	"time"

	"github.com/lourenci/github-alfred/lib/github"
	"github.com/lourenci/github-alfred/usecases/convertrepositorytoalfred"
	"github.com/stretchr/testify/require"
)

func TestGetUserReposInAlfred(t *testing.T) {
	t.Run("returns all user repos in alfred results", func(t *testing.T) {
		alfred := convertrepositorytoalfred.New(
			newFakeRepository(
				[]github.Repository{
					{
						Name:        "octocat/Hello-World",
						URL:         "https://github.com/octocat/Hello-World",
						Description: "This your first repo!",
						SshURL:      "git@github.com:octocat/Hello-World.git",
					},
					{
						Name:        "octocat/foo",
						URL:         "https://github.com/octocat/foo",
						Description: "foo repo",
						SshURL:      "git@github.com:octocat/foo.git",
					},
				},
			),
		).GetUserReposInAlfred(time.Hour * 1)

		require.Equal(
			t,
			convertrepositorytoalfred.Alfred{
				Cache: convertrepositorytoalfred.Cache{Seconds: 60 * 60},
				Items: []convertrepositorytoalfred.Item{
					{
						UID:      "octocat/Hello-World",
						Title:    "octocat/Hello-World",
						Subtitle: "This your first repo!",
						Match:    "octocat/Hello-World This your first repo!",
						Text:     convertrepositorytoalfred.Text{Copy: "git@github.com:octocat/Hello-World.git"},
						Mods: convertrepositorytoalfred.Mods{
							Cmd: convertrepositorytoalfred.Cmd{
								Subtitle: "⌘-C to copy git url",
								Arg:      "https://github.com/octocat/Hello-World",
							},
						},
						Arg: "https://github.com/octocat/Hello-World",
					},
					{
						UID:      "octocat/foo",
						Title:    "octocat/foo",
						Subtitle: "foo repo",
						Match:    "octocat/foo foo repo",
						Text:     convertrepositorytoalfred.Text{Copy: "git@github.com:octocat/foo.git"},
						Mods: convertrepositorytoalfred.Mods{
							Cmd: convertrepositorytoalfred.Cmd{
								Subtitle: "⌘-C to copy git url",
								Arg:      "https://github.com/octocat/foo",
							},
						},
						Arg: "https://github.com/octocat/foo",
					},
				},
			},
			alfred,
		)
	})

	t.Run("converts the provided cache to seconds", func(t *testing.T) {
		alfred := convertrepositorytoalfred.New(
			newFakeRepository(
				[]github.Repository{
					{
						Name:        "octocat/Hello-World",
						URL:         "https://github.com/octocat/Hello-World",
						Description: "This your first repo!",
						SshURL:      "git@github.com:octocat/Hello-World.git",
					},
					{
						Name:        "octocat/foo",
						URL:         "https://github.com/octocat/foo",
						Description: "foo repo",
						SshURL:      "git@github.com:octocat/foo.git",
					},
				},
			),
		)

		require.Equal(
			t,
			convertrepositorytoalfred.Alfred{
				Cache: convertrepositorytoalfred.Cache{Seconds: 7200},
			}.Cache,
			alfred.GetUserReposInAlfred(time.Hour*2).Cache,
		)
	})
}

type fakeRepository struct {
	repos []github.Repository
}

func newFakeRepository(repos []github.Repository) fakeRepository {
	return fakeRepository{repos: repos}

}

func (f fakeRepository) GetAllUserRepos() []github.Repository {
	return f.repos
}
