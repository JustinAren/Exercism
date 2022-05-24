package isogram

import "strings"

func IsIsogram(word string) bool {
	word = strings.ToLower(word)
	word = strings.ReplaceAll(word, "-", "")
	word = strings.ReplaceAll(word, " ", "")
	runeCount := make(map[rune]int)
	for _, r := range word {
		_, exists := runeCount[r]
		if exists {
			return false
		}

		runeCount[r] = 1
	}
	return true
}
