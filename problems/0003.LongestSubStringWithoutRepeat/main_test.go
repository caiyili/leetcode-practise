package leetcode

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLengthOfLongestSubString(t *testing.T) {
	tt := []struct {
		input   string
		exceped int
	}{
		{
			"abcabcbb",
			3,
		},
		{
			"bbbbbb",
			1,
		},
		{
			"pwwkew",
			3,
		},
	}

	assertion := assert.New(t)
	for _, c := range tt {
		got := lengthOfLongestSubstring(c.input)
		assertion.Equalf(c.exceped, got, "case#{%v}  got#{%v}", c, got)
	}
}
