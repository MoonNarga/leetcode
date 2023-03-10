package solution

func Search(nums []int, target int) int {
	for l, r := 0, len(nums)-1; l <= r; {
		mid := (l + r) / 2
		if nums[mid] == target {
			return mid
		}
		if nums[r] == target {
			return r
		}
		if nums[l] < nums[mid] {
			if target < nums[mid] {
				if target >= nums[l] {
					r = mid - 1
				} else {
					l = mid + 1
				}
			} else {
				l = mid + 1
			}
		} else {
			if target < nums[mid] {
				r = mid - 1
			} else {
				if target <= nums[r] {
					l = mid + 1
				} else {
					r = mid - 1
				}
			}
		}
	}
	return -1
}
