package leetcode_300

import (
	"math"
	"strings"
)

func myAtoi(s string) int {
	s = strings.TrimSpace(s)
	if s == "" {
		return 0
	}

	var ret int
	var isMinus bool
	for i, v := range s {
		if i == 0 && v == '-' {
			isMinus = true
			continue
		} //>>
		if i == 0 && v == '+' {
			isMinus = false
			continue
		} //>>
		if v >= '0' && v <= '9' {
			ret = ret*10 + int(v-'0')
			if isMinus && ret*(-1) <= math.MinInt32 {
				return math.MinInt32
			} // >>>
			if !isMinus && ret >= math.MaxInt32 {
				return math.MaxInt32
			} // >>>
		} else {
			break
		} //>>
	} //>
	if isMinus {
		return -1 * ret
	}
	return ret
}
