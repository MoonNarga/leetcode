package solution

func numOfZeors(v int) int {
	if v < 5 {
		return 0
	}
	res := 0
	for v != 0 {
		res += v / 5
		v /= 5
	}
	return res
}

func preimageSizeFZF(k int) int {
	i, j := 1, 10000000000
	for i <= j {
		mid := (i + j) / 2
		if zeors := numOfZeors(mid); zeors == k {
			return 5
		} else if zeors > k {
			j = mid - 1
		} else {
			i = mid + 1
		}
	}
	return 0
}
