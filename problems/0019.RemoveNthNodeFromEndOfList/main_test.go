package leetcode

import (
	"fmt"
	"testing"
)

func (node *ListNode) String() string {
	result := make([]byte, 1)
	for ; node != nil; node = node.Next {
		result = append(result, byte('0'+node.Val))
	}
	return string(result)
}

func buildList(arr []int) *ListNode {
	if len(arr) == 0 {
		return nil
	}
	head := &ListNode{arr[0], nil}
	current := head
	for i, n := range arr {
		if i > 0 {
			node := &ListNode{n, nil}
			current.Next = node
			current = current.Next
		}
	}
	return head
}

func ListEqual(node1, node2 *ListNode) bool {
	for node1 != nil || node2 != nil {
		if node1 == nil || node2 == nil || node1.Val != node2.Val {
			return false
		}
		node1 = node1.Next
		node2 = node2.Next
	}
	return true
}

var testCases = []struct {
	input   []int
	n       int
	exceped []int
}{
	{
		[]int{2, 4, 3, 1},
		4,
		[]int{4, 3, 1},
	},
	{
		[]int{2, 4, 3, 9},
		3,
		[]int{2, 3, 9},
	},
	{
		[]int{2, 4, 3},
		1,
		[]int{2, 4},
	},
}

func TestRemoveNthFromEnd(t *testing.T) {
	for _, tc := range testCases {
		got := removeNthFromEnd(buildList(tc.input), tc.n)
		fmt.Println(got)
		if !ListEqual(buildList(tc.exceped), got) {
			t.Errorf("test failed %+v, got:%v", tc, got)
		}
	}
}

func TestRemoveNthFromEndV2(t *testing.T) {
	for _, tc := range testCases {
		got := removeNthFromEndV2(buildList(tc.input), tc.n)
		fmt.Println(got)
		if !ListEqual(buildList(tc.exceped), got) {
			t.Errorf("test failed %+v, got:%v", tc, got)
		}
	}
}
