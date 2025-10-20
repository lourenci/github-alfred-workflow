package repository

import (
	"github.com/lourenci/github-alfred/lib/github"
	"github.com/lourenci/github-alfred/usecases/getopenpullsinalfred/vo"
)

type repository struct {
	httpClient github.GitHub
}

func New(githubHttpClient github.GitHub) repository {
	return repository{httpClient: githubHttpClient}
}

func (s repository) OpenPulls(repo vo.Repo, user string) []github.Pull {
	return s.httpClient.OpenPulls(string(repo), user)
}
