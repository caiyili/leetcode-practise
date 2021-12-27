package leetcode

import (
	"github.com/stretchr/testify/assert"
	"sort"
	"testing"
)

var testCases = []struct {
	input   int
	exceped []string
}{
	{
		1,
		[]string{"()"},
	},
	{
		2,
		[]string{"()()", "(())"},
	},
	{
		3,
		[]string{"()()()", "()(())", "(())()", "((()))", "(()())"},
	},
}

func TestGenerateParentheses(t *testing.T) {
	for _, c := range testCases {
		got := generateParenthesis(c.input)
		sort.Strings(c.exceped)
		sort.Strings(got)
		assert.Equal(t, c.exceped, got, c)
	}
}
