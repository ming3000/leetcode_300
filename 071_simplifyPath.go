package leetcode_300

import "strings"

func simplifyPath(path string) string {
	tempStack := make([]string, 0, len(path))
	items := strings.Split(path, "/")
	for _, v := range items {
		if v == "" || v == "." {
			continue
		} else if v == ".." {
			if len(tempStack) > 0 {
				tempStack = tempStack[:len(tempStack)-1]
			} //>>>
		} else {
			tempStack = append(tempStack, v)
		} //>>
	} //>
	return "/" + strings.Join(tempStack, "/")
}
