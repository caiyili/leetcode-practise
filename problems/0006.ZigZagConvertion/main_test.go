package leetcode

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestConvert(t *testing.T) {
	tt := []struct {
		input   string
		num     int
		exceped string
	}{
		{
			"ABC",
			3,
			"ABC",
		},
		{
			"PAYPALISHIRING",
			3,
			"PAHNAPLSIIGYIR",
		},
		{
			"PAYPALISHIRING",
			4,
			"PINALSIGYAHRPI",
		},
		{
			"PAYPALISHIRING",
			1,
			"PAYPALISHIRING",
		},
		{
			"A1B2C3",
			2,
			"ABC123",
		},
	}

	assertion := assert.New(t)
	for i, c := range tt {
		got := convert(c.input, c.num)
		assertion.Equalf(c.exceped, got, "case#%d#{%v}", i, c)
	}
}
