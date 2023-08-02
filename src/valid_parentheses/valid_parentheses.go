package validparentheses

// VALID PARENTHESES
// https://leetcode.com/problems/valid-parentheses/
// Given a string s containing just the characters '(', ')', '{', '}', '[' and ']',
// determine if the input string is valid.
//
// An input string is valid if:
//
// Open brackets must be closed by the same type of brackets.
// Open brackets must be closed in the correct order.
// Every close bracket has a corresponding open bracket of the same type.

func Solve(s string) bool {
	bracketPairs := map[rune]rune{
		'(': ')',
		'{': '}',
		'[': ']',
	}
	var search []rune

	for _, char := range s {
		if closing, ok := bracketPairs[char]; ok {
			// Char is Opening
			// append the closing rune to the search "stack"
			search = append(search, closing)
			continue
		}
		if len(search) == 0 || char != search[len(search)-1] {
			// there was no opened bracket, or the last opened bracket was of a different kind
			return false
		} else {
			// closing bracket matched the type of the last opened bracket
			// therefore it was valid and can be removed from the search slice
			search = search[:len(search)-1]
		}
	}
	return len(search) == 0
}
