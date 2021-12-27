package leetcode

import (
	"strings"
)

func intToRoman(num int) string {
	type baseConvert struct {
		Roman string
		Digit int
	}
	convertMap := []baseConvert{
		{"M", 1000},
		{"CM", 900},
		{"D", 500},
		{"CD", 400},
		{"C", 100},
		{"XC", 90},
		{"L", 50},
		{"XL", 40},
		{"X", 10},
		{"IX", 9},
		{"V", 5},
		{"IV", 4},
		{"I", 1},
	}
	var result string
	for _, base := range convertMap {
		if num >= base.Digit {
			times := num / base.Digit
			result += strings.Repeat(base.Roman, times)
			num = num - base.Digit*times
		}
	}
	return result
}
