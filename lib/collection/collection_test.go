package collection_test

import (
	"strconv"
	"testing"

	"github.com/lourenci/github-alfred/lib/collection"
	"github.com/stretchr/testify/require"
)

func TestMap(t *testing.T) {
	t.Run("returns a new slice with the results of applying the function to each item", func(t *testing.T) {
		require.Equal(
			t,
			[]string{},
			collection.Map(
				[]int{},
				func(it int) string {
					return strconv.Itoa(it)
				},
			),
		)

		require.Equal(
			t,
			[]string{"1", "2", "3", "4", "5"},
			collection.Map(
				[]int{1, 2, 3, 4, 5},
				func(it int) string {
					return strconv.Itoa(it)
				},
			),
		)

		require.Equal(
			t,
			[]int{1, 2, 3, 4, 5},
			collection.Map(
				[]int{1, 2, 3, 4, 5},
				func(it int) int {
					return it
				},
			),
		)
	})
}

func TestDedup(t *testing.T) {
	t.Run("returns a new slice with the duplicates removed", func(t *testing.T) {
		require.Equal(
			t,
			[]int{},
			collection.Dedup([]int{}),
		)

		require.Equal(
			t,
			[]int{1, 2, 3, 5},
			collection.Dedup([]int{1, 2, 3, 2, 5}),
		)

		type payload struct {
			text      string
			more_text string
		}
		require.Equal(
			t,
			[]payload{
				{
					text:      "a",
					more_text: "b",
				},
				{
					text:      "b",
					more_text: "a",
				},
			},
			collection.Dedup(
				[]payload{
					{
						text:      "a",
						more_text: "b",
					},
					{
						text:      "a",
						more_text: "b",
					},
					{
						text:      "b",
						more_text: "a",
					},
				},
			),
		)
	})
}
