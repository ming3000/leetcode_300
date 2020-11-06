package leetcode_300

import "strings"

func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	s = strings.ToLower(s)
	t = strings.ToLower(t)
	ca := [26]byte{}
	for _, v := range s {
		ca[v-'a'] += 1
	}
	for _, v := range t {
		ca[v-'a'] -= 1
	}
	for _, v := range ca {
		if v != 0 {
			return false
		} //>>
	} //>
	return true
}
