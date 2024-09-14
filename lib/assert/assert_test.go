package assert_test

import (
	"errors"
	"testing"

	"github.com/lourenci/github-alfred/lib/assert"
)

func TestNoError(t *testing.T) {
	t.Run("returns the result when there is no error", func(t *testing.T) {
		noErrorFn := func() (int, error) {
			return 1, nil
		}

		assert.Equals(t, assert.NoError(noErrorFn()), 1)
	})

	t.Run("panics with the error messages when there is an error", func(t *testing.T) {
		errorFn := func() (int, error) {
			return 0, errors.New("some error")
		}

		assert.PanicsWithMessage(
			t,
			func() { assert.NoError(errorFn()) },
			"some error",
		)
	})
}
