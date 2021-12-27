package leetcode

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"sort"
	"testing"
)

var testCases = []struct {
	input    []int
	expected []int
}{
	{
		[]int{1, 2, 3},
		[]int{1, 3, 2},
	},
	{
		[]int{3, 2, 1},
		[]int{1, 2, 3},
	},
	{
		[]int{1, 1, 5},
		[]int{1, 5, 1},
	},
	{
		[]int{8, 9, 6, 10, 7, 2},
		[]int{8, 9, 7, 2, 6, 10},
	},
}

func TestNextPermutation(t *testing.T) {
	for _, c := range testCases {
		input := make([]int, len(c.input))
		copy(input, c.input)
		fmt.Println("copy input:", input)
		nextPermutation(input)
		assert.Equal(t, c.expected, input, c)
	}
}

func sortInts(nums []int) {
	sort.Ints(nums)
}

func ExampleIntSlice() {
	input := []int{3, 1, 2, 4}
	sortInts(input)
	fmt.Println(input)
	//OUTPUT:
	//[1,2,3,4]
}
