package leetcode

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

var testCases = []struct {
	input   []int
	n       int
	exceped []int
}{
	{
		[]int{1, 1, 2},
		2,
		[]int{1, 2},
	},
}

func TestRemoveDuplicate(t *testing.T) {
	for _, c := range testCases {
		n := removeDuplicates(c.input)
		assert.Equal(t, c.n, n, c)
	}
}
