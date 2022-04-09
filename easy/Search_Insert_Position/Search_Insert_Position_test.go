package search_insert_position

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type Case struct {
	name   string
	nums   []int
	target int
	expect int
}

func TestSearchInsert(t *testing.T) {
	var cases = []Case{
		{"TestCase 1", []int{1, 3, 5, 6}, 5, 2},
		{"TestCase 2", []int{1, 3, 5, 6}, 2, 1},
		{"TestCase 3", []int{1, 3, 5, 6}, 7, 4},
		{"TestCase 4", []int{1, 3, 5, 6, 9, 10, 30, 50, 61, 100}, 69, 9},
	}

	ast := assert.New(t)

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			ast.Equal(
				c.expect,
				searchInsert(c.nums, c.target),
				"输入: %v, 目标： %v", c.nums, c.target,
			)
		})
	}
}
