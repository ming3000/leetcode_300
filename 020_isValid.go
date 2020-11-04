package leetcode_300

func isValid(s string) bool {
	cm := map[byte]byte{'(': ')', '[': ']', '{': '}'}
	cs := make([]byte, 0, len(s))

	for i := 0; i < len(s); i++ {
		if s[i] == '(' || s[i] == '[' || s[i] == '{' {
			cs = append(cs, s[i])
		} else {
			if len(cs) == 0 {
				return false
			} //>>>
			top := cs[len(cs)-1]
			cs = cs[:len(cs)-1]
			if cm[top] != s[i] {
				return false
			} //>>>
		} //>>
	} //>
	return len(cs) == 0
}
