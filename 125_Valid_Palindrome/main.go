package main

import (
	"regexp"
	"strings"
)

// nonAlphaNumericCharacter is a regular expression that matches any character
// that is NOT a letter, digit, or space. Used to clean input strings.
var nonAlphaNumericCharacter = regexp.MustCompile(`[^a-zA-Z0-9 ]+`)

// clearString removes all non-alphanumeric characters from the input string
// and converts it to lowercase. This helps in normalizing the string for
// palindrome checking.
func clearString(str string) string {
	str = nonAlphaNumericCharacter.ReplaceAllString(str, "")
	lowerS := strings.ToLower(str)
	return lowerS
}

// isPalindrome checks if a given string is a palindrome, ignoring case and
// non-alphanumeric characters. It uses two pointers (left and right) to
// compare characters from both ends, moving towards the center.
func isPalindrome(s string) bool {
	s = clearString(s)
	right := len(s) - 1
	left := 0
	for left < right {
		if s[left] != s[right] {
			// If characters at left and right pointers do not match,
			// the string is not a palindrome.
			return false
		} else {
			// Move both pointers towards the center.
			left++
			right--
		}
	}
	// If all characters matched, the string is a palindrome.
	return true
}

// main demonstrates the usage of clearString and isPalindrome functions.
// It tests a sample string, prints whether it is a palindrome, and shows
// the cleaned version of the string.
func main() {
	var str = "A!@#1B2b1@#@a"
	cleaned := clearString(str)
	if isPalindrome(str) {
		println("The string is a palindrome.")
		println("Cleaned string:", cleaned)
	} else {
		println("The string is not a palindrome.")
	}
}
