package solution

import "strconv"

func findNthDigit(n int) int {
	top := 10
	num := 9
	numsum := 0
	bit := 1
	for {
		if n <= numsum+bit*num {
			sub := n - numsum
			sub /= bit
			bottom := top/10 - 1
			bottom += sub
			numsum += sub * bit
			n -= numsum
			if n == 0 {
				return bottom % 10
			}
			str := strconv.Itoa(bottom + 1)
			b := str[n-1]
			res := int(b - '0')
			return res
		}
		numsum += num * bit
		top *= 10
		num = top - top/10
		bit++
	}
}
