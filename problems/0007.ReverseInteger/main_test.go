package leetcode

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestReverse(t *testing.T) {
	tt := []struct {
		input   int
		exceped int
	}{
		{123, 321},
		{1, 1},
		{-123, -321},
		{1534236469, 0},
	}
	for i, c := range tt {
		got := reverse(c.input)
		assert.Equalf(t, c.exceped, got, "case#%d#{%v} got:%v", i, c, got)
	}
}
