package collection_test

import (
	"strconv"
	"testing"

	"github.com/lourenci/github-alfred/lib/assert"
	"github.com/lourenci/github-alfred/lib/collection"
)

func TestMap(t *testing.T) {
	t.Run("returns a new slice with the results of applying the function to each item", func(t *testing.T) {
		assert.Equals(
			t,
			collection.Map(
				[]int{},
				func(it int) string {
					return strconv.Itoa(it)
				},
			),
			[]string{},
		)

		assert.Equals(
			t,
			collection.Map(
				[]int{1, 2, 3, 4, 5},
				func(it int) string {
					return strconv.Itoa(it)
				},
			),
			[]string{"1", "2", "3", "4", "5"},
		)

		assert.Equals(
			t,
			collection.Map(
				[]int{1, 2, 3, 4, 5},
				func(it int) int {
					return it
				},
			),
			[]int{1, 2, 3, 4, 5},
		)
	})
}

func TestDedup(t *testing.T) {
	t.Run("returns a new slice with the duplicates removed", func(t *testing.T) {
		assert.Equals(
			t,
			collection.Dedup([]int{}),
			[]int{},
		)

		assert.Equals(
			t,
			collection.Dedup([]int{1, 2, 3, 2, 5}),
			[]int{1, 2, 3, 5},
		)

		type payload struct {
			text      string
			more_text string
		}
		assert.Equals(
			t,
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
		)
	})
}
