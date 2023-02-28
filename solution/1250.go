package solution

func isGoodArray(nums []int) bool {
	g := 0
	for _, v := range nums {
		g = gcd(g, v)
		if g == 1 {
			return true
		}
	}
	return false
}

func gcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}
