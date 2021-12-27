package leetcode

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

var testCases = []struct {
	input     []int
	n         int
	expected  []int
	expectedN int
}{
	{
		[]int{3, 2, 2, 3},
		3,
		[]int{2, 2},
		2,
	},
	{
		[]int{1, 1, 2},
		2,
		[]int{1, 1},
		2,
	},
}

func TestRemoveDuplicate(t *testing.T) {
	for _, c := range testCases {
		n := removeElement(c.input, c.n)
		assert.Equalf(t, c.expectedN, n, "n=%d excepedN:%d", n, c.expectedN)
		assert.Equal(t, c.expected, c.input[:n])
	}
}
