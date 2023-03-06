package solution

func printBin(num float64) string {
	base := 0.5
	n := 0.0
	res := "0."
	for n < num && len(res) <= 32 {
		if n + base <= num {
			n += base
			res += "1"
		} else {
			res += "0"
		}
		base *= 0.5
	}
	if n != num {
		return "ERROR"
	}
	return res
}
