package house_robber

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
	{"TestCase 1", []int{1, 2, 3, 1}, 4},
	{"TestCase 2", []int{2, 7, 9, 3, 1}, 12},
	{"TestCase 3", []int{0}, 0},
	{"TestCase 4", []int{0, 8, 12, 9, 1, 0, 13}, 30},
	{"TestCase 5", []int{1, 2}, 2},
	{"TestCase 6", []int{9, 2, 1}, 10},
}

func TestRob(t *testing.T) {
	ast := assert.New(t)

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			ast.Equal(c.expect, rob(c.input), "输入: %v", c.input)
		})
	}
}
