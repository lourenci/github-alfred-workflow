package getopenpullsinalfred

import (
	"fmt"
	"strings"
)

type Repo string

func MustParse(name string) Repo {
	if !strings.Contains(name, "/") {
		panic(fmt.Errorf(`invalid repo name: "%s"`, name))
	}

	return Repo(name)
}
