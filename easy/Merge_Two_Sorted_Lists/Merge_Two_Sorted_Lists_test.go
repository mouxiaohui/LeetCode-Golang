package merge_two_sorted_lists

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type Case struct {
	name   string
	input  input
	expect *ListNode
}

type input struct {
	l1 *ListNode
	l2 *ListNode
}

var cases = []Case{
	{
		"TesCase 1",
		input{l1: generateListNode([]int{1, 2, 4}), l2: generateListNode([]int{1, 3, 4})},
		generateListNode([]int{1, 1, 2, 3, 4, 4}),
	},
	{
		"TesCase 2",
		input{l1: nil, l2: nil},
		nil,
	},
	{
		"TesCase 3",
		input{l1: nil, l2: &ListNode{Val: 0}},
		&ListNode{Val: 0},
	},
}

func TestNumberOfLines(t *testing.T) {
	ast := assert.New(t)

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			ast.Equal(c.expect, mergeTwoLists(c.input.l1, c.input.l2), "%s", c.name)
		})
	}
}

func generateListNode(nums []int) *ListNode {
	var res *ListNode
	for _, num := range nums {
		if res == nil {
			res = &ListNode{Val: num}
		} else {
			res.Append(num)
		}
	}

	return res
}
