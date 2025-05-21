package assert_test

import (
	"errors"
	"testing"

	"github.com/lourenci/github-alfred/lib/assert"
	"github.com/stretchr/testify/require"
)

func TestNoError(t *testing.T) {
	t.Run("returns the result when there is no error", func(t *testing.T) {
		noErrorFn := func() (int, error) {
			return 1, nil
		}

		require.Equal(t, 1, assert.NoError(noErrorFn()))
	})

	t.Run("panics with the error messages when there is an error", func(t *testing.T) {
		errorFn := func() (int, error) {
			return 0, errors.New("some error")
		}

		require.PanicsWithError(
			t,
			"some error",
			func() { assert.NoError(errorFn()) },
		)
	})
}
