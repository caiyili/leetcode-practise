package leetcode

type ListNode struct {
	Val  int
	Next *ListNode
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	var nodeTable = make([]*ListNode, 0)
	var index = 0
	for p := head; p != nil; p = p.Next {
		nodeTable = append(nodeTable, p)
		index++
	}
	length := index
	nodeNum := length - n
	if nodeNum == 0 {
		return head.Next
	} else {
		preNode := nodeTable[nodeNum-1]
		preNode.Next = nodeTable[nodeNum].Next
	}
	return head
}

func removeNthFromEndV2(head *ListNode, n int) *ListNode {
	var root = &ListNode{}
	root.Next = head
	var moveCount = 0
	var fastPtr, slowPtr = head, root
	for fastPtr.Next != nil {
		moveCount++
		fastPtr = fastPtr.Next
		if moveCount >= n {
			slowPtr = slowPtr.Next
		}
	}
	slowPtr.Next = slowPtr.Next.Next
	return root.Next
}
