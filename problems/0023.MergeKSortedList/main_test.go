package leetcode

import (
	"fmt"
	"testing"
)

func (node *ListNode) String() string {
	result := make([]int, 1)
	for ; node != nil; node = node.Next {
		result = append(result, node.Val)
	}
	return fmt.Sprintf("%v", result)
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
	arr1, arr2 []int
	exceped    []int
}{
	{
		[]int{1, 2, 4, 31},
		[]int{5, 6, 44},
		[]int{1, 2, 4, 5, 6, 31, 44},
	},
	{
		[]int{2, 3, 9},
		[]int{5, 6, 8},
		[]int{2, 3, 5, 6, 8, 9},
	},
	{
		[]int{2, 14, 23},
		[]int{5, 15, 24},
		[]int{2, 5, 14, 15, 23, 24},
	},
}

func TestMergeKList(t *testing.T) {
	for _, tc := range testCases {
		node1 := buildList(tc.arr1)
		fmt.Println(node1)
		node2 := buildList(tc.arr2)
		fmt.Println(node2)
		exceped := buildList(tc.exceped)
		fmt.Println(exceped)
		got := mergeKLists([]*ListNode{node1, node2})
		fmt.Println(got)
		if !ListEqual(exceped, got) {
			t.Errorf("test failed %v, got:%v", tc, got)
		}
	}
}

func TestMergeKListV2(t *testing.T) {
	for _, tc := range testCases {
		node1 := buildList(tc.arr1)
		fmt.Println(node1)
		node2 := buildList(tc.arr2)
		fmt.Println(node2)
		exceped := buildList(tc.exceped)
		fmt.Println(exceped)
		got := mergeKListsV2([]*ListNode{node1, node2})
		fmt.Println(got)
		if !ListEqual(exceped, got) {
			t.Errorf("test failed %v, got:%v", tc, got)
		}
	}
}
