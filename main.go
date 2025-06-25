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
	"github.com/lourenci/github-alfred/usecases/convertrepositorytoalfred"
	"github.com/lourenci/github-alfred/usecases/convertrepositorytoalfred/repository"
)

func main() {
	token := os.Args[1]
	cacheInMinutes, _ := strconv.Atoi(os.Args[2])

	useCase := convertrepositorytoalfred.New(
		repository.New(github.New(token, http.New())),
	)

	json := assert.NoError(json.Marshal(useCase.GetUserReposInAlfred(
		assert.NoError(time.ParseDuration(fmt.Sprintf("%dm", cacheInMinutes))),
	)))

	fmt.Println(string(json))
}
