package leetcode_300

import "strings"

func isPalindrome125(s string) bool {
	s = strings.TrimSpace(s)
	s = strings.ToLower(s)
	left, right := 0, len(s)-1
	for left < right {
		if !isValidChar(s[left]) {
			left++
			continue
		} //>>
		if !isValidChar(s[right]) {
			right--
			continue
		} //>>
		if s[left] != s[right] {
			return false
		} //>>
		left++
		right--
	} //>
	return true
}

func isValidChar(ch byte) bool {
	if (ch >= '0' && ch <= '9') ||
		(ch >= 'a' && ch <= 'z') ||
		(ch >= 'A' && ch <= 'Z') {
		return true
	}
	return false
}
