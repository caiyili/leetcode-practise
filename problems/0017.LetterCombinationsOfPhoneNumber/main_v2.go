package leetcode

func letterCombinationsV2(digits string) []string {

	var letterMap = map[byte]string{
		'2': "abc",
		'3': "def",
		'4': "ghi",
		'5': "jkl",
		'6': "mno",
		'7': "pqrs",
		'8': "tuv",
		'9': "wxyz",
	}
	var lenDigit = len(digits)
	if lenDigit == 0 {
		return []string{}
	}
	var result = make([]string, 0)

	for i := 0; i < lenDigit; i++ {
		digit := digits[i]
		expandChar := letterMap[digit]
		if i == 0 {
			for _, c := range expandChar {
				result = append(result, string(c))
			}
		} else {
			newResult := make([]string, 0)
			for _, c := range expandChar {
				for _, r := range result {
					newResult = append(newResult, r+string(c))
				}
			}
			result = newResult
		}
	}

	return result
}
