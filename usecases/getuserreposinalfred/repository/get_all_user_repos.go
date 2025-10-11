package repository

import (
	"github.com/lourenci/github-alfred/lib/collection"
	"github.com/lourenci/github-alfred/lib/github"
)

type repository struct {
	httpClient github.GitHub
}

func New(githubHttpClient github.GitHub) repository {
	return repository{httpClient: githubHttpClient}
}

func (r repository) GetAllUserRepos() []github.Repository {
	channels := make(chan []github.Repository)
	func(c chan []github.Repository) {
		go func() {
			c <- r.httpClient.StarredRepos()
		}()
		go func() {
			c <- r.httpClient.UserRepos()
		}()
		go func() {
			c <- r.httpClient.WatchedRepos()
		}()
	}(channels)

	repos := collection.Dedup(append(append(<-channels, <-channels...), <-channels...))

	return repos
}
