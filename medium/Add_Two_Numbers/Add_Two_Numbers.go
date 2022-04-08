package add_two_numbers

// 给你两个 非空 的链表，表示两个非负的整数。它们每位数字都是按照 逆序 的方式存储的，并且每个节点只能存储 一位 数字。
// 请你将两个数相加，并以相同形式返回一个表示和的链表。
// 你可以假设除了数字 0 之外，这两个数都不会以 0 开头。

// 每个链表中的节点数在范围 [1, 100] 内
// 0 <= Node.val <= 9
// 题目数据保证列表表示的数字不含前导零

// Example 1:
// Input: l1 = [2,4,3], l2 = [5,6,4]
// Output: [7,0,8]
// Explanation: 342 + 465 = 807.

// Example 2:
// Input: l1 = [0], l2 = [0]
// Output: [0]

// Example 3:
// Input: l1 = [9,9,9,9,9,9,9], l2 = [9,9,9,9]
// Output: [8,9,9,9,0,0,0,1]

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

// 为测试准备
func (h *ListNode) Append(i int) {
	for h.Next != nil {
		h = h.Next
	}
	h.Next = &ListNode{Val: i}
}

// 为测试准备
func (h *ListNode) Show() []int {
	res := []int{h.Val}
	for h.Next != nil {
		h = h.Next
		res = append(res, h.Val)
	}

	return res
}

// 题解
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var res, node *ListNode
	addOne := false

	for l1 != nil || l2 != nil {
		a, b := 0, 0
		if l1 != nil {
			a = l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			b = l2.Val
			l2 = l2.Next
		}

		sum := a + b
		if addOne {
			sum = sum + 1
			addOne = false
		}
		if sum >= 10 {
			sum = sum - 10
			addOne = true
		}

		if res == nil {
			res = &ListNode{Val: sum}
			node = res
		} else {
			node.Next = &ListNode{Val: sum}
			node = node.Next
		}
	}
	if addOne {
		node.Next = &ListNode{Val: 1}
	}

	return res
}
