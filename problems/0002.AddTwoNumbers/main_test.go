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

func TestAddTwoNumbers(t *testing.T) {
	tt := []struct {
		arr1, arr2 []int
		exceped    []int
	}{
		{
			[]int{2, 4, 3, 1},
			[]int{5, 6, 4},
			[]int{7, 0, 8, 1},
		},
		{
			[]int{2, 4, 3, 9},
			[]int{5, 6, 8},
			[]int{7, 0, 2, 0, 1},
		},
		{
			[]int{2, 4, 3},
			[]int{5, 6, 4},
			[]int{7, 0, 8},
		},
	}

	for _, tc := range tt {
		node1 := buildList(tc.arr1)
		fmt.Println(node1)
		node2 := buildList(tc.arr2)
		fmt.Println(node2)
		exceped := buildList(tc.exceped)
		got := addTwoNumbers(node1, node2)
		fmt.Println(got)
		if !ListEqual(exceped, got) {
			t.Errorf("test failed %v, got:%v", tc, got)
		}
	}
}
