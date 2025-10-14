package getopenpullsinalfred_test

import (
	"testing"

	"github.com/lourenci/github-alfred/usecases/getopenpullsinalfred"
	"github.com/stretchr/testify/require"
)

func TestMustParse(t *testing.T) {
	t.Run("parses the Repo", func(t *testing.T) {
		require.Equal(
			t,
			getopenpullsinalfred.Repo("octocat/Hello-World"),
			getopenpullsinalfred.MustParseRepo("octocat/Hello-World"),
		)
	})

	t.Run("panics when name is invalid", func(t *testing.T) {
		require.PanicsWithError(
			t,
			`invalid repo name: ""`,
			func() {
				getopenpullsinalfred.MustParseRepo("")
			},
		)
		require.PanicsWithError(
			t,
			`invalid repo name: "invalid-repo"`,
			func() {
				getopenpullsinalfred.MustParseRepo("invalid-repo")
			},
		)
	})
}
