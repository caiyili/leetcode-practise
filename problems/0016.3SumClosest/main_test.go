package leetcode

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

var testCases = []struct {
	nums    []int
	target  int
	exceped int
}{
	{
		nums:    []int{-1, 2, 1, -4},
		target:  1,
		exceped: 2,
	},
}

func TestThreeeSumClosest(t *testing.T) {
	for _, c := range testCases {
		assert.Equal(t, c.exceped, threeSumClosest(c.nums, c.target), c)
	}
}
