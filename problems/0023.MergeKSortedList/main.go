package leetcode

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeKLists(lists []*ListNode) *ListNode {
	var root = &ListNode{}
	var k = len(lists)
	var currNode = make([]*ListNode, 0)
	for _, node := range lists {
		currNode = append(currNode, node)
	}
	var preNode = root
	for {
		minNodeIdx := -1
		nilCount := 0
		for i, node := range currNode {
			if node == nil {
				nilCount++
			} else if minNodeIdx < 0 || node.Val < currNode[minNodeIdx].Val {
				minNodeIdx = i
			}
		}
		if nilCount == k {
			break
		}
		newNode := currNode[minNodeIdx]
		currNode[minNodeIdx] = currNode[minNodeIdx].Next
		preNode.Next = newNode
		preNode = newNode
	}

	return root.Next
}

func mergeKListsV2(lists []*ListNode) *ListNode {
	length := len(lists)
	if length == 0 {
		return nil
	}
	if length == 1 {
		return lists[0]
	}
	mid := length / 2
	left := mergeKListsV2(lists[0:mid])
	right := mergeKListsV2(lists[mid:])
	return mergeTwoLists(left, right)
}

func mergeTwoLists(l1, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}
	if l1.Val < l2.Val {
		l1.Next = mergeTwoLists(l1.Next, l2)
		return l1
	} else {
		l2.Next = mergeTwoLists(l1, l2.Next)
		return l2
	}
}
