package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/lourenci/github-alfred/lib/assert"
	"github.com/lourenci/github-alfred/lib/github"
	"github.com/lourenci/github-alfred/lib/http"
	"github.com/lourenci/github-alfred/usecases/getopenpullsinalfred"
	getOpenPullsRepository "github.com/lourenci/github-alfred/usecases/getopenpullsinalfred/repository"
	"github.com/lourenci/github-alfred/usecases/getopenpullsinalfred/vo"
	"github.com/lourenci/github-alfred/usecases/getuserreposinalfred"
	getUsersRepository "github.com/lourenci/github-alfred/usecases/getuserreposinalfred/repository"
)

func main() {
	token := os.Args[1]

	switch os.Args[2] {
	case "repos":
		cacheInMinutes, _ := strconv.Atoi(os.Args[3])

		useCase := getuserreposinalfred.New(
			getUsersRepository.New(github.New(token, http.New())),
		)

		json := assert.NoError(json.Marshal(useCase.GetUserReposInAlfred(
			assert.NoError(time.ParseDuration(fmt.Sprintf("%dm", cacheInMinutes))),
		)))

		fmt.Println(string(json))
	case "pulls":
		repoName := os.Args[3]
		user := os.Args[4]

		useCase := getopenpullsinalfred.New(
			getOpenPullsRepository.New(github.New(token, http.New())),
		)

		json := assert.NoError(json.Marshal(useCase.GetUserOpenPullsOfRepo(
			vo.MustParseRepo(repoName),
			user,
		)))

		fmt.Println(string(json))
	default:
		panic("invalid option")
	}

}
