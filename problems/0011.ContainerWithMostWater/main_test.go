package leetcode

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

var testCases = []struct {
	input   []int
	exceped int
}{
	{
		[]int{1, 8, 6, 2, 5, 4, 8, 3, 7},
		49,
	},
	{
		[]int{2, 3, 4, 5, 18, 17, 6},
		17,
	},
}

func TestMaxAreaV1(t *testing.T) {
	for _, c := range testCases {
		assert.Equal(t, c.exceped, maxAreaV1(c.input), c)
	}
}

func TestMaxAreaV2(t *testing.T) {
	for _, c := range testCases {
		assert.Equal(t, c.exceped, maxArea(c.input), c)
	}
}
