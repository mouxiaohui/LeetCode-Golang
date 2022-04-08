package add_two_numbers

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
		"TestCase 1",
		input{
			l1: generateListNode([]int{2, 4, 3}),
			l2: generateListNode([]int{5, 6, 4}),
		},
		generateListNode([]int{7, 0, 8}),
	},
	{
		"TestCase 2",
		input{
			l1: generateListNode([]int{0}),
			l2: generateListNode([]int{0}),
		},
		generateListNode([]int{0}),
	},
	{
		"TestCase 3",
		input{
			l1: generateListNode([]int{9, 9, 9, 9, 9, 9, 9}),
			l2: generateListNode([]int{9, 9, 9, 9}),
		},
		generateListNode([]int{8, 9, 9, 9, 0, 0, 0, 1}),
	},
}

func TestAddTwoNumbers(t *testing.T) {
	ast := assert.New(t)

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			ast.Equal(
				c.expect,
				addTwoNumbers(c.input.l1, c.input.l2),
				"输入: l1 = %v, l2 = %v", c.input.l1.Show(), c.input.l2.Show(),
			)
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
