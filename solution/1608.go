package solution

import "sort"

func specialArray(nums []int) int {
	sort.Ints(nums)
	for i := len(nums); i > 0; i-- {
		if nums[len(nums)-i] >= i && (i == len(nums) || nums[len(nums)-i-1] < i) {
			return i
		}
	}
	return -1
}
