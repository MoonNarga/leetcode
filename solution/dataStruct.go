package solution

import (
	"sort"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type ListNode struct {
	Val  int
	Next *ListNode
}

type MaxHeapInt struct {
	sort.IntSlice
}

func (h MaxHeapInt) Less(i, j int) bool  { return h.IntSlice[i] > h.IntSlice[j] }
func (h *MaxHeapInt) Push(v interface{}) { h.IntSlice = append(h.IntSlice, v.(int)) }
func (h *MaxHeapInt) Pop() interface{} {
	v := h.IntSlice[len(h.IntSlice)-1]
	h.IntSlice = h.IntSlice[:len(h.IntSlice)-1]
	return v
}

type MinHeapInt struct {
	sort.IntSlice
}

func (h MinHeapInt) Less(i, j int) bool  { return h.IntSlice[i] < h.IntSlice[j] }
func (h *MinHeapInt) Push(v interface{}) { h.IntSlice = append(h.IntSlice, v.(int)) }
func (h *MinHeapInt) Pop() interface{} {
	v := h.IntSlice[len(h.IntSlice)-1]
	h.IntSlice = h.IntSlice[:len(h.IntSlice)-1]
	return v
}
