package leetcode

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

func letterCombinations(digits string) []string {
	return combination([]byte(digits))
}

func combination(digits []byte) []string {
	if len(digits) == 0 {
		return []string{}
	}
	first := digits[0]
	expand := letterMap[first]
	var result = make([]string, 0)
	if len(digits) == 1 {
		for i, _ := range expand {
			result = append(result, string(expand[i]))
		}
	} else {
		leftResult := combination(digits[1:])
		for _, char := range expand {
			for _, r := range leftResult {
				result = append(result, string(char)+r)
			}
		}
	}
	return result
}
