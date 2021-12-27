package leetcode

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIntToRoman(t *testing.T) {
	testCases := []struct {
		input   int
		exceped string
	}{
		{3, "III"},
		{4, "IV"},
		{9, "IX"},
		{58, "LVIII"},
		{1994, "MCMXCIV"},
	}

	for _, c := range testCases {
		assert.Equal(t, c.exceped, intToRoman(c.input), c)
	}
}
