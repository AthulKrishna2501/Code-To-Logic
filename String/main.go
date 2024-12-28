package main

import (
	"fmt"
	"strings"
	"unicode"
)

func longestString(s string) string {
	word := strings.Fields(s)
	longest := ""
	for _, s := range word {
		if len(s) > len(longest) {
			longest = s
		}
	}
	return longest
}

func checkSubString(s, sub string) bool {
	return strings.Contains(s, sub)
}

func isPalindrome(s string) bool {
	lowerString := strings.ToLower(s)
	var runes []rune

	for _, n := range lowerString {
		if unicode.IsLetter(n) || unicode.IsDigit(n) {
			runes = append(runes, n)
		}
	}

	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		if runes[i] != runes[j] {
			return false
		}
	}
	return true

}

func main() {
	newStr := ""
	str := "World"
	w := 'd'

	for _, n := range str {
		if n != w {
			newStr += string(n)
		}
	}
	fmt.Println(newStr)
}
