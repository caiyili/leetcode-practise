package leetcode

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestTwoSum(t *testing.T) {
	tt := []struct {
		nums    []int
		target  int
		exceped []int
	}{
		{
			[]int{0, 1, 2, 3, 5, 6},
			1,
			[]int{0, 1},
		},
		{
			[]int{0, 1, 2, 3, 5, 6},
			0,
			[]int{0, 0},
		},
		{
			[]int{0, 1, 2, 3, 5, 6},
			11,
			[]int{4, 5},
		},
		{
			[]int{0, 1, 3, 3, 5, 6},
			6,
			[]int{2, 3},
		},
	}

	assertion := assert.New(t)
	for _, t := range tt {
		got := findTwoSum(t.nums, t.target)
		assertion.Equalf(t.exceped, got, "nums:%v target=%v", t.nums, t.target)
	}
}
