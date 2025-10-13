package getopenpullsinalfred

import (
	"fmt"

	"github.com/lourenci/github-alfred/lib/collection"
	"github.com/lourenci/github-alfred/lib/github"
)

type Alfred struct {
	Items []Item `json:"items"`
}

type Item struct {
	UID      string `json:"uid"`
	Title    string `json:"title"`
	Subtitle string `json:"subtitle"`
	Match    string `json:"match"`
	Arg      string `json:"arg"`
}

type GithubRepository interface {
	UserOpenPullsOfRepo(repo, user string) []github.Pull
}

type UseCase struct {
	repository GithubRepository
}

func New(repository GithubRepository) UseCase {
	return UseCase{repository: repository}
}

func (r UseCase) GetUserReposInAlfred(repo, user string) Alfred {
	repositories := r.repository.UserOpenPullsOfRepo(repo, user)

	return Alfred{
		Items: collection.Map(repositories, func(repo github.Pull) Item {
			return Item{
				UID:      repo.URL,
				Title:    repo.Title,
				Subtitle: fmt.Sprintf("Created at %s", repo.CreatedAt),
				Match:    repo.Title,
				Arg:      repo.URL,
			}
		}),
	}
}
