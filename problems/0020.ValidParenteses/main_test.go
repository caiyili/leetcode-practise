package leetcode

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

var testCases = []struct {
	input   string
	exceped bool
}{
	{"({})", true},
	{"})", false},
	{"{})", false},
	{"", true},
}

func TestIsValid(t *testing.T) {
	for _, c := range testCases {
		assert.Equal(t, c.exceped, isValid(c.input), c)
	}
}
