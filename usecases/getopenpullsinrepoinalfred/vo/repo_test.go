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
