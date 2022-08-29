package solution

func shuffle(nums []int, n int) []int {
	var index int
	for i := 0; i < len(nums); i++ {
		if nums[i] < 0 {
			continue
		}
		index = i
		for nums[i] > 0 {
			if index < n {
				index = 2 * index
			} else {
				index = 2*(index-n) + 1
			}
			if index == i {
				nums[i] = -nums[i]
			} else {
				nums[i], nums[index] = nums[index], -nums[i]
			}
		}
	}
	for i, v := range nums {
		nums[i] = -v
	}
	return nums
}
