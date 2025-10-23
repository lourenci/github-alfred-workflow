package repository

import "github.com/lourenci/github-alfred/lib/github"

type repository struct {
	githubClient github.GitHub
}

func New(githubClient github.GitHub) repository {
	return repository{
		githubClient: githubClient,
	}
}

func (r repository) OpenPulls(user string) []github.Pull {
	return r.githubClient.OpenPulls(github.NewOpenPullsQuery(
		github.MustParseUserQuery(user),
	))
}
