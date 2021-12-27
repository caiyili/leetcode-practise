package leetcode

func romanToInt(roman string) int {
	convertMap := map[string]int{
		"CM": 900,
		"CD": 400,
		"XC": 90,
		"XL": 40,
		"IX": 9,
		"IV": 4,
		"M":  1000,
		"D":  500,
		"C":  100,
		"L":  50,
		"X":  10,
		"V":  5,
		"I":  1,
	}
	var result int
	length := len(roman)
	for i := 0; i < length; i++ {
		single := string(roman[i])
		pair := ""
		if i < length-1 {
			pair = string(roman[i : i+2])
		}
		if digit, ok := convertMap[pair]; ok {
			result += digit
			i++
		} else {
			digit = convertMap[single]
			result += digit
		}
	}
	return result
}
