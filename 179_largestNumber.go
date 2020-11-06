package leetcode_300

import (
	"sort"
	"strconv"
	"strings"
)

func largestNumber(nums []int) string {
	strNums := make([]string, len(nums))
	for i, v := range nums {
		strNums[i] = strconv.Itoa(v)
	}

	sort.Slice(strNums, func(i, j int) bool {
		return strNums[i]+strNums[j] > strNums[j]+strNums[i]
	})
	ret := strings.Join(strNums, "")
	if ret[0] == '0' {
		return "0"
	}
	return ret
}
