package house_robber_3

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type Case struct {
	name   string
	root   *TreeNode
	expect int
}

var cases = []Case{
	{
		"TestCase 1",
		&TreeNode{
			Val:   3,
			Left:  &TreeNode{Val: 2, Right: &TreeNode{Val: 3}},
			Right: &TreeNode{Val: 3, Right: &TreeNode{Val: 1}},
		},
		7,
	},
	{
		"TestCase 2",
		&TreeNode{
			Val:   3,
			Left:  &TreeNode{Val: 4, Left: &TreeNode{Val: 1}, Right: &TreeNode{Val: 3}},
			Right: &TreeNode{Val: 5, Right: &TreeNode{Val: 1}},
		},
		9,
	},
	{
		"TestCase 3",
		&TreeNode{
			Val:  4,
			Left: &TreeNode{Val: 1, Left: &TreeNode{Val: 2, Left: &TreeNode{Val: 3}}},
		},
		7,
	},
}

func TestRob3(t *testing.T) {
	ast := assert.New(t)

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			ast.Equal(c.expect, rob(c.root), c.name)
		})
	}
}
