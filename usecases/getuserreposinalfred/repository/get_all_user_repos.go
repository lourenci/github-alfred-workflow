package repository

import (
	"github.com/lourenci/github-alfred/lib/collection"
	"github.com/lourenci/github-alfred/lib/github"
)

func GetAllUserRepos(githubRepository github.GitHub) []github.Repository {
	channels := make(chan []github.Repository)
	func(c chan []github.Repository) {
		go func() {
			c <- githubRepository.StarredRepos()
		}()
		go func() {
			c <- githubRepository.UserRepos()
		}()
		go func() {
			c <- githubRepository.WatchedRepos()
		}()
	}(channels)

	repos := collection.Dedup(append(append(<-channels, <-channels...), <-channels...))

	return repos
}
