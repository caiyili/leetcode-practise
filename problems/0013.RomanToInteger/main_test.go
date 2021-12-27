package leetcode

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRomanToInt(t *testing.T) {
	testCases := []struct {
		input   string
		exceped int
	}{
		{"III", 3},
		{"IV", 4},
		{"IX", 9},
		{"LVIII", 58},
		{"MCMXCIV", 1994},
	}

	for _, c := range testCases {
		assert.Equal(t, c.exceped, romanToInt(c.input), c)
	}
}
