package main

import (
	"fmt"
)

func main() {
	keysPressed := "spuda"
	releaseTimes := []int{12,23,36,46,62}
	fmt.Printf("%c\n", slowestKey(releaseTimes, keysPressed))
}
