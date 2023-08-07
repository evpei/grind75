package validpalindrome

import "unicode"

// A phrase is a palindrome if, after converting all uppercase letters into lowercase letters
// and removing all non-alphanumeric characters, it reads the same forward and backward.
// Alphanumeric characters include letters and numbers.
//
// Given a string s, return true if it is a palindrome, or false otherwise.

func Solve(s string) bool {
	runes := []rune(s)
	startIndex := 0
	endIndex := len(runes) - 1

	for startIndex <= endIndex {
		if !isAlphanumeric(runes[startIndex]) {
			startIndex++
			continue
		} else if !isAlphanumeric(runes[endIndex]) {
			endIndex--
			continue
		}
		if unicode.ToLower(runes[startIndex]) != unicode.ToLower(runes[endIndex]) {
			return false
		}
		startIndex++
		endIndex--
	}

	return true
}

func isAlphanumeric(c rune) bool {
	return (c >= 'a' && c <= 'z') || (c >= 'A' && c <= 'Z') || (c >= '0' && c <= '9')
}
