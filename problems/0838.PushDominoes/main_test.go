package leetcode

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPushDominoes(t *testing.T) {
	var testCases = []struct {
		input    string
		expected string
	}{
		{
			"",
			"",
		},
		{
			"L",
			"L",
		},
		{
			"R",
			"R",
		},
		{
			"R.",
			"RR",
		},
		{
			"...",
			"...",
		},
		{
			"LRRLLRR...L...",
			"LRRLLRRR.LL...",
		},
		{
			"RR.L",
			"RR.L",
		},
		{
			".L.R...LR..L..",
			"LL.RR.LLRRLL..",
		},
	}
	for _, c := range testCases {
		assert.Equal(t, c.expected, pushDominoes(c.input), c)
	}
}
