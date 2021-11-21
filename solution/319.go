package solution

func bulbSwitch(n int) int {
	if n == 0 {
		return 0
	}
	if n == 1 {
		return n
	}
	bulb := make([]bool, n+1)
	for i := 2; i <= n; i++ {
		for j := i; j <= n; j += i {
			if bulb[j] == false {
				bulb[j] = true
			} else {
				bulb[j] = false
			}
		}
	}
	c := 0
	for i := 1; i <= n; i++ {
		if bulb[i] == false {
			c++
		}
	}
	return c
}
