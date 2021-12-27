package leetcode

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

var testCases = []struct {
	s        string
	words    []string
	expected []int
}{
	{
		s:        "barfoothefoobarman",
		words:    []string{"foo", "bar"},
		expected: []int{0, 9},
	},
	{
		s:        "wordgoodgoodgoodbestword",
		words:    []string{"word", "good", "best", "word"},
		expected: []int{},
	},
	{
		s:        "wordgoodgoodgoodbestword",
		words:    []string{"word", "good", "best", "good"},
		expected: []int{8},
	},
}

func TestFindSubstring(t *testing.T) {
	for _, c := range testCases {
		assert.Equal(t, c.expected, findSubstring(c.s, c.words), c)
	}
}
