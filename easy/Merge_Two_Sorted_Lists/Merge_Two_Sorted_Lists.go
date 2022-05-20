package merge_two_sorted_lists

/*
第 21 题

将两个升序链表合并为一个新的 升序 链表并返回。新链表是通过拼接给定的两个链表的所有节点组成的。

示例 1：
输入：l1 = [1,2,4], l2 = [1,3,4]
输出：[1,1,2,3,4,4]

示例 2：
输入：l1 = [], l2 = []
输出：[]

示例 3：
输入：l1 = [], l2 = [0]
输出：[0]

提示：
两个链表的节点数目范围是 [0, 50]
-100 <= Node.val <= 100
l1 和 l2 均按 非递减顺序 排列
*/

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

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	dummy := &ListNode{}
	n, n1, n2 := dummy, list1, list2

	for n1 != nil && n2 != nil {
		if n1.Val < n2.Val {
			n.Next = n1
			n1 = n1.Next
		} else {
			n.Next = n2
			n2 = n2.Next
		}

		n = n.Next
	}

	if n1 != nil {
		n.Next = n1
	}

	if n2 != nil {
		n.Next = n2
	}

	return dummy.Next
}
