package leetcode_300

import "strings"

func reverseWords(s string) string {
	s = strings.TrimSpace(s)
	strList := strings.Split(s, " ")
	retList := make([]string, 0, len(strList))
	for i := len(strList) - 1; i >= 0; i-- {
		if strList[i] == "" {
			continue
		}
		retList = append(retList, strList[i])
	}
	return strings.Join(retList, " ")
}
