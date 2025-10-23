package vo

import (
	"fmt"
	"regexp"
	"strings"
)

type Repo string

func MustParseRepo(name string) Repo {
	if !strings.Contains(name, "/") {
		panic(fmt.Errorf(`invalid repo name: "%s"`, name))
	}

	return Repo(name)
}

func MustParseRepoFromUrl(url string) Repo {
	matches := regexp.MustCompile("/repos/([^/]+/[^/]+)").FindStringSubmatch(url)

	if len(matches) == 0 {
		panic(fmt.Errorf(`invalid URL: "%s"`, url))
	}

	return MustParseRepo(matches[1])
}
