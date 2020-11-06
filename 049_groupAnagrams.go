package leetcode_300

import (
	"sort"
	"strings"
)

func groupAnagrams(strs []string) [][]string {
	keyMap := make(map[string][]string)
	for _, str := range strs {
		key := generateKey(str)
		if _, ok := keyMap[key]; !ok {
			keyMap[key] = make([]string, 0)
		} //>>
		keyMap[key] = append(keyMap[key], str)
	} //>

	ret := make([][]string, 0)
	for key := range keyMap {
		ret = append(ret, keyMap[key])
	}
	return ret
}

func generateKey(s string) string {
	s = strings.TrimSpace(s)
	s = strings.ToLower(s)
	b := []byte(s)
	sort.Slice(b, func(i, j int) bool {
		return b[i] < b[j]
	})
	return string(b)
}
