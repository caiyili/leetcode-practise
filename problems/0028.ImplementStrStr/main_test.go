package leetcode

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

var testCases = []struct {
	haystack string
	needle   string
	expected int
}{
	{
		"hello",
		"ll",
		2,
	},
	{
		"aaaaaa",
		"bba",
		-1,
	},
	{
		"",
		"",
		0,
	},
	{
		"aaa",
		"a",
		0,
	},
}

func TestStrstr(t *testing.T) {
	for _, c := range testCases {
		assert.Equal(t, c.expected, strstr(c.haystack, c.needle), c)
	}
}
