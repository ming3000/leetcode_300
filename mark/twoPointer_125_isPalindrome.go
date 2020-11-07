package mark

import "strings"

func isPalindrome(s string) bool {
	s = strings.TrimSpace(s)
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
	if ch >= 'a' && ch <= 'z' {
		return true
	}
	if ch >= 'A' && ch <= 'Z' {
		return true
	}
	if ch >= '0' && ch <= '9' {
		return true
	}
	return false
}
