package solution

import (
	"container/heap"
	"sort"
)

func mincostToHireWorkers(quality []int, wage []int, k int) float64 {
	indices := make([]int, len(quality))
	for i := range indices {
		indices[i] = i
	}
	sort.Slice(indices, func(i, j int) bool {
		return wage[indices[i]]*quality[indices[j]] < wage[indices[j]]*quality[indices[i]]
	})
	h := &MaxHeapInt{}
	sumQuality := 0
	for i := 0; i < k; i++ {
		sumQuality += quality[indices[i]]
		heap.Push(h, quality[indices[i]])
	}
	res := float64(sumQuality) * (float64(wage[indices[k-1]]) / float64(quality[indices[k-1]]))
	for i := k; i < len(indices); i++ {
		v := heap.Pop(h).(int)
		sumQuality -= v
		sumQuality += quality[indices[i]]
		heap.Push(h, quality[indices[i]])
		res = Min(res, float64(sumQuality)*(float64(wage[indices[i]])/float64(quality[indices[i]])))
	}
	return res
}
