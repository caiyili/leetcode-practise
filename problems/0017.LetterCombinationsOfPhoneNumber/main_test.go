package leetcode

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

var testCases = []struct {
	input   string
	exceped []string
}{
	{
		input:   "23",
		exceped: []string{"ad", "ae", "af", "bd", "be", "bf", "cd", "ce", "cf"},
	},
}

func TestLetterCombination(t *testing.T) {
	for _, c := range testCases {
		assert.Equal(t, c.exceped, letterCombinations(c.input), c)
	}
}

func TestLetterCombinationV2(t *testing.T) {
	for _, c := range testCases {
		assert.Equal(t, c.exceped, letterCombinationsV2(c.input), c)
	}
}
