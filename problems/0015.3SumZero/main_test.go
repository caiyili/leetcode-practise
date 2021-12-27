package leetcode

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestThreeSum(t *testing.T) {
	testCases := []struct {
		input   []int
		exceped [][]int
	}{
		{
			input: []int{-1, 0, 1, 2, -1, -4},
			exceped: [][]int{
				[]int{-1, -1, 2},
				[]int{-1, 0, 1},
			},
		},
	}

	for _, c := range testCases {
		assert.Equal(t, c.exceped, threeSum(c.input), c)
	}
}

func TestThreeSumV2(t *testing.T) {
	testCases := []struct {
		input   []int
		exceped [][]int
	}{
		{
			input: []int{-1, 0, 1, 2, -1, -4},
			exceped: [][]int{
				[]int{-1, -1, 2},
				[]int{-1, 0, 1},
			},
		},
		{
			input: []int{0, 0, 0, 0},
			exceped: [][]int{
				[]int{0, 0, 0},
			},
		},
		{
			input: []int{-2, 0, 1, 1, 2},
			exceped: [][]int{
				[]int{-2, 0, 2},
				[]int{-2, 1, 1},
			},
		},
	}

	for _, c := range testCases {
		assert.Equal(t, c.exceped, threeSumV2(c.input), c)
	}
}
