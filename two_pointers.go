package main

import "strings"

// 125 Valid Palindrome
func IsPalindrome(s string) bool {
	if len(s) == 0 {
		return true
	}
	isAlphanumeric := func(char byte) bool {
		return (char >= 'a' && char <= 'z') ||
			(char >= 'A' && char <= 'Z') ||
			(char >= '0' && char <= '9')
	}
	l, r := 0, len(s)-1
	for l < r {
		for !isAlphanumeric(s[l]) && l < r {
			l++
		}
		for !isAlphanumeric(s[r]) && l < r {
			r--
		}
		if strings.ToLower(string(s[l])) != strings.ToLower(string(s[r])) {
			return false
		}
		l++
		r--
	}
	return true
}
