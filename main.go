package main

import (
	"fmt"
	"strconv"
)

func isAdditiveNumber(num string) bool {
	numArray := []byte(num)
	lmax := len(num) / 3
	var tempNum1, tempNum2, tempNum3, tempsum int
	for l1 := 1; l1 <= lmax; l1++ {
		tempNum1, _ = strconv.Atoi(string(numArray[0:l1]))
		for l2 := 1; l2 <= lmax; l2++ {
			tempNum2, _ = strconv.Atoi(string(numArray[l1 : l1+l1]))
			tempsum = tempNum1 + tempNum2
		}
	}
	return false
}

func main() {
	num := "112358"
	fmt.Println(isAdditiveNumber(num))
}
