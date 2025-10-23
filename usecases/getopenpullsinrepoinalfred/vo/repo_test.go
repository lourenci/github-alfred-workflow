package vo_test

import (
	"testing"

	"github.com/lourenci/github-alfred/usecases/getopenpullsinrepoinalfred/vo"
	"github.com/stretchr/testify/require"
)

func TestMustParse(t *testing.T) {
	t.Run("parses the Repo", func(t *testing.T) {
		require.Equal(
			t,
			vo.Repo("octocat/Hello-World"),
			vo.MustParseRepo("octocat/Hello-World"),
		)
	})

	t.Run("panics when name is invalid", func(t *testing.T) {
		require.PanicsWithError(
			t,
			`invalid repo name: ""`,
			func() {
				vo.MustParseRepo("")
			},
		)
		require.PanicsWithError(
			t,
			`invalid repo name: "invalid-repo"`,
			func() {
				vo.MustParseRepo("invalid-repo")
			},
		)
	})
}

func TestMustParseRepoFromUrl(t *testing.T) {
	t.Run("panics when URL is invalid", func(t *testing.T) {
		require.PanicsWithError(
			t,
			`invalid URL: ""`,
			func() {
				vo.MustParseRepoFromUrl("")
			},
		)
		require.PanicsWithError(
			t,
			`invalid URL: "https://api.github.com/batterseapower/pinyin-toolkit"`,
			func() {
				vo.MustParseRepoFromUrl("https://api.github.com/batterseapower/pinyin-toolkit")
			},
		)
	})

	t.Run("parses the Repo from URL", func(t *testing.T) {
		require.Equal(
			t,
			vo.Repo("batterseapower/pinyin-toolkit"),
			vo.MustParseRepoFromUrl("https://api.github.com/repos/batterseapower/pinyin-toolkit"),
		)
	})
}
