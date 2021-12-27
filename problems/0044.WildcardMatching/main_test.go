package leetcode

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

var testCases = []struct {
	s, p     string
	expected bool
}{
	{"aa", "a", false},
	{"aa", "*", true},
	{"cb", "?a", false},
	{"", "", true},
	{"", "****", true},
	{"adceb", "*a*b", true},
	{"acdcd", "a*c?b", false},
}

func TestIsMatch(t *testing.T) {
	for _, c := range testCases {
		assert.Equal(t, c.expected, isMatch(c.s, c.p))
	}
}

func TestCompress(t *testing.T) {
	tc := []struct {
		input, expected string
	}{
		{"***", "*"},
		{"abc", "abc"},
		{"***abc", "*abc"},
		{"a***b", "a*b"},
	}
	for _, c := range tc {
		assert.Equal(t, c.expected, compress(c.input))
	}
}
