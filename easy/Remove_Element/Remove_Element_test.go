package remove_element

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type Case struct {
	name   string
	input  input
	expect int
}

type input struct {
	nums []int
	val  int
}

var cases = []Case{
	{
		"TesCase 1",
		input{
			[]int{0, 1, 2, 2, 3, 0, 4, 2},
			2,
		},
		5,
	},
	{
		"TesCase 2",
		input{
			[]int{3, 2, 2, 3},
			3,
		},
		2,
	},
	{
		"TesCase 3",
		input{
			[]int{0},
			0,
		},
		0,
	},
	{
		"TesCase 4",
		input{
			[]int{9, 9, 9, 8, 3, 2, 1},
			9,
		},
		4,
	},
}

func TestRemoveElement(t *testing.T) {
	ast := assert.New(t)

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			ast.Equal(c.expect, removeElement(c.input.nums, c.input.val), "%s", c.name)
		})
	}
}
