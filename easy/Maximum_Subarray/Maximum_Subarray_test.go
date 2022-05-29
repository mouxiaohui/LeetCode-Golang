package maximum_subarray

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
		[]int{-2, 1, -3, 4, -1, 2, 1, -5, 4},
		6,
	},
	{
		"TesCase 2",
		[]int{1},
		1,
	},
	{
		"TesCase 3",
		[]int{-2, -1},
		-1,
	},
	{
		"TesCase 4",
		[]int{-2, 1},
		1,
	},
}

func TestMaxSubArray(t *testing.T) {
	ast := assert.New(t)

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			ast.Equal(c.expect, maxSubArray(c.input), "%s", c.name)
		})
	}
}
