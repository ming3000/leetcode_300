package mark

import "strings"

func isAnagram(s string, t string) bool {
	s = strings.TrimSpace(s)
	s = strings.ToLower(s)
	t = strings.TrimSpace(t)
	t = strings.ToLower(t)
	if len(s) != len(t) {
		return false
	}

	checkArr := [26]byte{}
	for _, v := range s {
		checkArr[v-'a'] += 1
	}
	for _, v := range t {
		checkArr[v-'a'] -= 1
	}

	for _, v := range checkArr {
		if v != 0 {
			return false
		} //>>
	} //>
	return true
}
