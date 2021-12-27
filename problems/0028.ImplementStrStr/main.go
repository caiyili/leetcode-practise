package leetcode

func strstr(haystack string, needle string) int {
	var haystackLength = len(haystack)
	var needleLength = len(needle)
	if needleLength == 0 {
		return 0
	}
	var pos = -1
	for i := 0; i < haystackLength; i++ {
		if i+needleLength > haystackLength {
			break
		}
		substr := haystack[i : i+needleLength]
		if substr == needle {
			pos = i
			break
		}
	}
	return pos
}
