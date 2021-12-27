package leetcode

func findSubstring(s string, words []string) []int {
	if len(words) == 0 {
		return []int{}
	}
	totalLength := len(s)
	wordLength := len(words[0])
	wordCount := len(words)
	totalWordLength := wordLength * wordCount

	var result = make([]int, 0)
	for index := 0; index < totalLength; index++ {
		if index+totalWordLength > totalLength {
			break
		}
		wordContainer := make(map[string]int)
		for _, w := range words {
			wordContainer[w]++
		}
		for j := index; j < index+totalWordLength; j += wordLength {
			splitWord := s[j : j+wordLength]
			if _, ok := wordContainer[splitWord]; ok {
				wordContainer[splitWord]--
				if wordContainer[splitWord] == 0 {
					delete(wordContainer, splitWord)
				}
			} else {
				break
			}
		}
		if len(wordContainer) == 0 {
			result = append(result, index)
		}
	}
	return result
}
