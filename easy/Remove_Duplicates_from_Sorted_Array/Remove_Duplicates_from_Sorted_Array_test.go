package remove_duplicates_from_sorted_array

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type Case struct {
	name   string
	input  []int
	expect int
}

var cases = []Case{
	{
		"TesCase 1",
		[]int{1, 1, 2},
		2,
	},
	{
		"TesCase 2",
		[]int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4},
		5,
	},
	{
		"TesCase 3",
		[]int{0, 1, 2, 2, 3, 4, 4},
		5,
	},
	{
		"TesCase 4",
		[]int{2, 2, 2, 2},
		1,
	},
}

func TestRemoveDuplicates(t *testing.T) {
	ast := assert.New(t)

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			ast.Equal(c.expect, removeDuplicates(c.input), "%s", c.name)
		})
	}
}
