package luhn

import (
	"strings"
	"unicode"
)

func Valid(id string) bool {
	id = Reverse(strings.ReplaceAll(id, " ", ""))
	if len(id) < 2 {
		return false
	}
	resultSlice := make([]int, len(id))
	for i, r := range id {
		if !unicode.IsDigit(r) {
			return false
		}
		n := int(r) - 48
		if i%2 == 1 {
			n *= 2
			if n > 9 {
				n -= 9
			}
		}
		resultSlice[i] = n
	}
	sum := 0
	for i := 0; i < len(resultSlice); i++ {
		sum += resultSlice[i]
	}
	return sum%10 == 0
}

func Reverse(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}
