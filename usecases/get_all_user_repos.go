package usecases

import "github.com/lourenci/github-alfred/lib/github"

func GetAllUserRepos(githubRepository github.GitHub) []github.Repository {
	var repos []github.Repository

	repos = append(
		append(
			append(
				repos,
				githubRepository.StarredRepos()...,
			),
			githubRepository.UserRepos()...,
		),
		githubRepository.WatchedRepos()...,
	)

	return repos
}
