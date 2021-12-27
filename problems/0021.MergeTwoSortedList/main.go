package leetcode

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	var root = &ListNode{}
	var p1, p2 = l1, l2
	var curr = root
	for p1 != nil || p2 != nil {
		var newNode *ListNode
		if p1 == nil {
			newNode = p2
			p2 = p2.Next
		} else if p2 == nil {
			newNode = p1
			p1 = p1.Next
		} else {
			if p1.Val <= p2.Val {
				newNode = p1
				p1 = p1.Next
			} else {
				newNode = p2
				p2 = p2.Next
			}
		}
		curr.Next = newNode
		curr = curr.Next
	}
	return root.Next
}
