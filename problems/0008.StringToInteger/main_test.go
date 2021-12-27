package leetcode

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMyAtoi(t *testing.T) {
	tt := []struct {
		input   string
		exceped int
	}{
		{"123", 123},
		{"-123", -123},
		{" 123", 123},
		{"   -123", -123},
		{"a23", 0},
		{"12a3", 12},
		{"99999999999999999", 1<<31 - 1},
		{"-99999999999999999999", -(1 << 31)},
		{"9223372036854775808", 2147483647},
	}

	for i, c := range tt {
		got := myAtoi(c.input)
		assert.Equalf(t, c.exceped, got, "case#%d#{%v}", i, c)
	}
}
