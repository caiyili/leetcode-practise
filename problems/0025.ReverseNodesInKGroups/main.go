package leetcode

import (
	"fmt"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func displayListNode(node *ListNode) {
	for p := node; p != nil; p = p.Next {
		fmt.Printf("->%d", p.Val)
	}
	fmt.Println("")
}

func reverseKGroup(head *ListNode, k int) *ListNode {
	if k < 2 {
		return head
	}
	var root = &ListNode{0, head}
	var groupFirstPre = root
	for {
		var groupLast = groupFirstPre
		var count = 0
		for groupLast.Next != nil && count < k {
			groupLast = groupLast.Next
			count++
		}
		if count != k {
			break
		}

		groupLastNext := groupLast.Next
		//需要对preFirst到last之间的节点转换
		first := groupFirstPre.Next
		pp := first
		p := pp.Next
		for p != groupLastNext {
			pNext := p.Next
			p.Next = pp

			pp = p
			p = pNext
		}
		groupFirstPre.Next = groupLast
		first.Next = groupLastNext

		groupFirstPre = first
	}

	return root.Next
}
