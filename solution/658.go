package solution

import "sort"

func absDiff(a, b int) int {
	if a < b {
		return b - a
	}
	return a - b
}

func findClosestElements(arr []int, k int, x int) []int {
	l, r := 0, len(arr)-1
	if arr[len(arr)-1] <= x {
		l = len(arr) - 1
	} else if arr[0] >= x {
		l = 0
	} else {
		for l < r-1 {
			mid := (l + r) / 2
			diff := arr[mid] - x
			if diff == 0 {
				l = mid
				break
			}
			if diff > 0 {
				r = mid
			} else {
				l = mid
			}
		}
	}
	cnt := 0
	r = l + 1
	res := make([]int, k)
	for cnt < k {
		if l < 0 {
			res[cnt] = arr[r]
			r++
		} else if r >= len(arr) {
			res[cnt] = arr[l]
			l--
		} else if absDiff(arr[l], x) <= absDiff(arr[r], x) {
			res[cnt] = arr[l]
			l--
		} else {
			res[cnt] = arr[r]
			r++
		}
		cnt++
	}
	sort.Ints(res)
	return res
}
