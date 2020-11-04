package leetcode_300

func generateParenthesis(n int) []string {
	ret := make([]string, 0)
	dfsGenerateParenthesis(0, 0, n, "", &ret)
	return ret
}

func dfsGenerateParenthesis(lc, rc int, n int, path string, ret *[]string) {
	if lc == n && rc == n {
		*ret = append(*ret, path)
		return
	}

	if lc < n {
		dfsGenerateParenthesis(lc+1, rc, n, path+"(", ret)
	}
	if rc < n && rc < lc {
		dfsGenerateParenthesis(lc, rc+1, n, path+")", ret)
	}
}
