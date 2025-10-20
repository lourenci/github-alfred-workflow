package getuserreposinalfred

import (
	"fmt"
	"time"

	"github.com/lourenci/github-alfred/lib/collection"
	"github.com/lourenci/github-alfred/lib/github"
)

type Alfred struct {
	Cache Cache  `json:"cache"`
	Items []Item `json:"items"`
}

type Cache struct {
	Seconds int `json:"seconds"`
}

type Item struct {
	UID      string `json:"uid"`
	Title    string `json:"title"`
	Subtitle string `json:"subtitle"`
	Match    string `json:"match"`
	Text     Text   `json:"text"`
	Mods     Mods   `json:"mods"`
	Arg      string `json:"arg"`
}

type Text struct {
	Copy string `json:"copy"`
}

type Mods struct {
	Cmd Cmd `json:"cmd"`
	Alt Alt `json:"alt"`
}

type Cmd struct {
	Subtitle string `json:"subtitle"`
	Arg      string `json:"arg"`
}

type Alt struct {
	Subtitle string `json:"subtitle"`
	Arg      string `json:"arg"`
}

type GithubRepository interface {
	GetAllUserRepos() []github.Repository
}

type UseCase struct {
	repository GithubRepository
}

func New(repository GithubRepository) UseCase {
	return UseCase{repository: repository}
}

func (r UseCase) GetUserReposInAlfred(cacheDuration time.Duration) Alfred {
	repositories := r.repository.GetAllUserRepos()

	return Alfred{
		Cache: Cache{Seconds: int(cacheDuration.Seconds())},
		Items: collection.Map(repositories, func(repo github.Repository) Item {
			return Item{
				UID:      repo.Name,
				Title:    repo.Name,
				Subtitle: repo.Description,
				Match:    fmt.Sprintf("%s %s", repo.Name, repo.Description),
				Text:     Text{Copy: repo.SshURL},
				Mods: Mods{
					Cmd: Cmd{
						Subtitle: "⌘-C to copy git url | ⌘-return to open in browser",
						Arg:      repo.URL,
					},
					Alt: Alt{
						Subtitle: "See options",
						Arg:      fmt.Sprintf("%s,%s", repo.URL, repo.Name),
					},
				},
				Arg: repo.URL,
			}
		}),
	}
}
