package mark

func search081(nums []int, target int) bool {
	left, right := 0, len(nums)-1
	for left <= right {
		mid := (left + right) / 2
		if nums[mid] == target {
			return true
		} //>>

		if nums[left] == nums[mid] {
			left++
		} else if nums[left] < nums[mid] {
			if nums[left] <= target && nums[mid] > target {
				right = mid - 1
			} else {
				left = mid + 1
			} //>>>
		} else {
			if nums[right] >= target && nums[mid] < target {
				left = mid + 1
			} else {
				right = mid - 1
			} //>>>
		} //>>
	} //>
	return false
}

//
//func search081(nums []int, target int) bool {
//	if len(nums) == 0 {
//		return false
//	}
//
//	low, pivot, high := 0, findPivot081(nums), len(nums)-1
//	if pivot == 0 || nums[low] > target {
//		low, high = pivot, len(nums)-1
//	} else {
//		low, high = 0, pivot-1
//	}
//	for low <= high {
//		mid := (low + high) / 2
//		if nums[mid] >= target {
//			high = mid-1
//		} else {
//			low = mid + 1
//		} //>>
//	} //>
//	if low >= len(nums) || nums[low] != target {
//		return false
//	}
//	return true
//}
//
//func findPivot081(nums []int) int {
//	low, high := 0, len(nums)-1
//	if nums[low] <= nums[high] {
//		return low
//	}
//	for low <= high {
//		mid := (low + high) / 2
//		if nums[mid] > nums[mid+1] {
//			return mid + 1
//		} else {
//			if nums[low] < nums[mid] {
//				low = mid + 1
//			} else {
//				high = mid
//			} //>>>
//		} //>>
//	} //>
//	return 0
//}
