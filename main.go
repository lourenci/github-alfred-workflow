package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strconv"

	"github.com/lourenci/github-alfred/lib/assert"
	"github.com/lourenci/github-alfred/lib/collection"
	"github.com/lourenci/github-alfred/lib/github"
	"github.com/lourenci/github-alfred/lib/http"
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
}

type Cmd struct {
	Subtitle string `json:"subtitle"`
	Arg      string `json:"arg"`
}

func main() {
	token := os.Args[1]
	cacheInMinutes, _ := strconv.Atoi(os.Args[2])

	githubApi := github.New(token, http.New())

	channels := make(chan []github.Repository)
	func(c chan []github.Repository) {
		go func() {
			c <- githubApi.StarredRepos()
		}()
		go func() {
			c <- githubApi.UserRepos()
		}()
		go func() {
			c <- githubApi.WatchedRepos()
		}()
	}(channels)

	repos := collection.Dedup(append(append(<-channels, <-channels...), <-channels...))

	alfred := Alfred{
		Cache: Cache{Seconds: cacheInMinutes * 60},
		Items: collection.Map(repos, func(repo github.Repository) Item {
			return Item{
				UID:      repo.Name,
				Title:    repo.Name,
				Subtitle: repo.Description,
				Match:    fmt.Sprintf("%s %s", repo.Name, repo.Description),
				Text:     Text{Copy: repo.SshURL},
				Mods: Mods{
					Cmd: Cmd{
						Subtitle: "⌘-C to copy git url",
						Arg:      repo.URL,
					},
				},
				Arg: repo.URL,
			}
		}),
	}

	json := assert.NoError(json.Marshal(alfred))
	fmt.Println(string(json))
}
