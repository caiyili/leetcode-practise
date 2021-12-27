package leetcode

type ListNode struct {
	Val  int
	Next *ListNode
}

func swapPairs(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	var root = &ListNode{0, head}
	var preFirst = root
	var nextNode *ListNode
	second := preFirst.Next.Next
	for second != nil {
		nextNode = second.Next
		second.Next = preFirst.Next
		preFirst.Next = second
		second.Next.Next = nextNode

		preFirst = second.Next
		if nextNode == nil {
			break
		}
		second = nextNode.Next
	}
	return root.Next
}
