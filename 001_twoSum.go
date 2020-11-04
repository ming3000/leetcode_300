package leetcode_300

func twoSum(nums []int, target int) []int {
	if len(nums) < 2 {
		return nil
	}
	cm := make(map[int]int)
	for i, v := range nums {
		diff := target - v
		if index, ok := cm[diff]; ok {
			return []int{i, index}
		} // if>>
		cm[v] = i
	} // for>
	return nil
}
