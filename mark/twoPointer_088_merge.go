package mark

func merge(nums1 []int, m int, nums2 []int, n int) {
	retIndex, i, j := m+n-1, m-1, n-1
	for i >= 0 && j >= 0 {
		if nums1[i] < nums2[j] {
			nums1[retIndex] = nums2[j]
			j--
		} else {
			nums1[retIndex] = nums1[i]
			i--
		} //>>
		retIndex--
	} //>
	for j >= 0 {
		nums1[retIndex] = nums2[j]
		j--
		retIndex--
	}
}
