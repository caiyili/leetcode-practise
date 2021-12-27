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
	input   []int
	exceped []int
}{
	{
		[]int{1, 2, 4, 31},
		[]int{2, 1, 31, 4},
	},
	{
		[]int{2, 3, 9},
		[]int{3, 2, 9},
	},
	{
		[]int{5, 15, 24, 1, 19, 20, 10},
		[]int{15, 5, 1, 24, 20, 19, 10},
	},
	{
		[]int{10},
		[]int{10},
	},
	{
		[]int{},
		[]int{},
	},
}

func TestSwapPairs(t *testing.T) {
	for _, c := range testCases {
		input := buildList(c.input)
		exceped := buildList(c.exceped)
		got := swapPairs(input)
		if !ListEqual(exceped, got) {
			t.Errorf("case#%v got:%v", c, got)
		}
	}
}
