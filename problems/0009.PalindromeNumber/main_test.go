package leetcode

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIsPalindrome(t *testing.T) {
	tt := []struct {
		num     int
		exceped bool
	}{
		{1, true},
		{12, false},
		{-100, false},
		{1001, true},
		{1234321, true},
	}

	for _, c := range tt {
		assert.Equal(t, c.exceped, isPalindrome(c.num))
	}
}
