package getopenpullsinalfred

import "errors"

type Repo string

func MustParse(name string) Repo {
	if name == "" {
		panic(errors.New("invalid repo name: \"\""))
	}

	return Repo(name)
}
