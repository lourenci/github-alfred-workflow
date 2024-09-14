package assert

import (
	"reflect"
	"testing"
)

func Equals[T any](t *testing.T, actual T, expected T) {
	t.Helper()

	// `DeepEqual` considers empty slice/arrays as not equals.
	if reflect.TypeOf(actual).Kind() == reflect.Slice {
		if reflect.ValueOf(actual).Len() == 0 && reflect.ValueOf(expected).Len() == 0 {
			return
		}
	}

	if !reflect.DeepEqual(actual, expected) {
		t.Fatalf(`expected: "%v"; got it: "%v"`, expected, actual)
	}
}

func PanicsWithMessage(t *testing.T, fn func(), msg string) {
	t.Helper()

	defer func() {
		r := recover()

		if r == nil {
			t.Fatalf("expected to panic, but it did not")
		}
		if r.(error).Error() != msg {
			t.Fatalf(
				`expected to panic with the message "%s", but it panicked with the message "%s"`,
				msg,
				r,
			)
		}
	}()

	fn()
}

func NotPanics(t *testing.T, fn func()) {
	t.Helper()

	defer func() {
		r := recover()

		if r != nil {
			t.Fatalf("expected not to panic, but it panicked")
		}
	}()

	fn()
}

func NoError[T any](result T, err error) T {
	if err != nil {
		panic(err)
	}

	return result
}
