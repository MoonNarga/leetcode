package solution

import "container/list"

type node struct {
	in  int
	out []int
}

func sequenceReconstruction(nums []int, sequences [][]int) bool {
	index := make(map[int]int)
	for i, v := range nums {
		index[v] = i
	}
	nodeList := make([]node, len(nums))
	for _, v := range nodeList {
		v.in = 0
		v.out = make([]int, 0)
	}
	for _, v := range sequences {
		l := len(v)
		for i := 0; i < l-1; i++ {
			if index[v[i]] >= index[v[i+1]] {
				return false
			}
			nodeList[v[i]-1].out = append(nodeList[v[i]-1].out, v[i+1]-1)
			nodeList[v[i+1]-1].in++
		}
	}
	queue := list.New()
	for _, v := range nodeList {
		if v.in == 0 {
			queue.PushBack(v)
		}
	}
	for queue.Len() != 0 {
		if queue.Len() > 1 {
			return false
		}
		temp := queue.Front().Value.(node)
		for _, v := range temp.out {
			nodeList[v].in--
			if nodeList[v].in == 0 {
				queue.PushBack(nodeList[v])
			}
		}
		queue.Remove(queue.Front())
	}
	return true
}
