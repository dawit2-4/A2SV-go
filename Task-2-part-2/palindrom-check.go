package main

import (
	"fmt"
	"regexp"
	"strings"
)

// IsPalindrome checks if a string is a palindrome (ignoring punctuation, spaces, and case)
func IsPalindrome(s string) bool {
	// Convert to lowercase
	s = strings.ToLower(s)

	// Remove all digit characters
	reg, _ := regexp.Compile(`[^\w]`)
	s = reg.ReplaceAllString(s, "")

	// Compare the string with its reverse
	n := len(s)
	for i := 0; i < n/2; i++ {
		if s[i] != s[n-1-i] {
			return false
		}
	}
	return true
}

func main() {
	// Example tests
	tests := []string{
		"Madam",
		"A man, a plan, a canal: Panama",
		"Hello, World!",
		"Was it a car or a cat I saw?",
		"racecar",
	}

	for _, test := range tests {
		fmt.Printf("Is \"%s\" a palindrome? %v\n", test, IsPalindrome(test))
	}
}
