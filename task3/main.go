// Что такое палиндром? Как сравнивать строку с её обратной версией?
// Напиши функцию, проверяющую, является ли строка палиндромом:
package task3
import (
	"strings"
	"unicode"
)

func IsPalindrome(s string) bool {
	cleaned := strings.Map(func(r rune) rune {
		if unicode.IsLetter(r) {
			return unicode.ToLower(r)
		}
		return -1
	}, s)

	for i := range len(cleaned)/2 {
		if cleaned[i] != cleaned[len(cleaned)-1-i] {
			return false
		}
	}
	return true
}

func main() {
	examples := []string{
		"madam",
		"racecar",
		"12321",
		"hello",
		"12345",
	}

	for _, example := range examples {
		if IsPalindrome(example) {
			println(example)
		} else {
			println(false)
		}
	}
}