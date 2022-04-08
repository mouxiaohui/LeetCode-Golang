package two_sum

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type Case struct {
	name   string
	nums   []int
	target int
	expect []int
}

func TestTwoSum(t *testing.T) {
	var cases = []Case{
		{"TestCase 1", []int{2, 7, 11, 15}, 9, []int{0, 1}},
		{"TestCase 2", []int{3, 2, 4}, 6, []int{1, 2}},
		{"TestCase 3", []int{3, 3}, 6, []int{0, 1}},
	}

	ast := assert.New(t)

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			ast.Equal(c.expect, twoSum(c.nums, c.target), "输入:%v", c.nums)
		})
	}

}
