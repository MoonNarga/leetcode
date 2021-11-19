package main

import (
	"fmt"
)

func integerReplacement(n int) int {
	step := 0
	for {
		if n == 1 {
			return step
		}
		if n == 3 {
			return step + 2
		}
		if n&3 == 3 {
			n++
		} else if n&3 == 1 {
			n--
		} else {
			n /= 2
		}
		step += 1
	}
}

func main() {
	var n int
	fmt.Scan(&n)
	fmt.Println(integerReplacement(n))
}
