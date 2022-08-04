package solution

import "sort"

func minSubsequence(nums []int) []int {
	sort.Slice(nums, func(i, j int) bool { return nums[i] > nums[j] })
	tailSum, headSum := 0, 0
	for _, v := range nums {
		tailSum += v
	}
	index := 0
	for i, v := range nums {
		headSum += v
		tailSum -= v
		if headSum > tailSum {
			index = i
			break
		}
	}
	return nums[:index+1]
}
