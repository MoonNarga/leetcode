package solution

import (
	"strconv"
	"strings"
)

func recu(n, i int, res []int, logs []string, execStack [][]string, subTime []int) int {
	length := len(logs)
	log := strings.Split(logs[i], ":")
	time := 0
	top := 0
	execStack[top] = log
	top++
	i++
	for top != 0 && i < length {
		log = strings.Split(logs[i], ":")
		execStack[top] = log
		top++
		i++
		if execStack[top-1][1] == "start" {
			continue
		}
		time2, _ := strconv.Atoi(log[2])
		time1, _ := strconv.Atoi(execStack[top-2][2])
		time = time2 - time1 + 1
		subTime[top-2] += time
		event, _ := strconv.Atoi(log[0])
		res[event] += time - subTime[top-1]
		subTime[top-1] = 0
		top -= 2
	}
	step := i
	return step
}

func exclusiveTime(n int, logs []string) []int {
	res := make([]int, n)
	length := len(logs)
	execStack := make([][]string, length)
	subTime := make([]int, length)
	for i := 0; i < length; {
		step := recu(n, i, res, logs, execStack, subTime)
		i = step
	}
	return res
}
