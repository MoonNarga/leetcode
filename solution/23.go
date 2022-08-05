package solution

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
import (
	"container/heap"
)

func mergeKLists(lists []*ListNode) *ListNode {
	h := make(ListNodeHeap, 0)
	for _, v := range lists {
		if v != nil {
			heap.Push(&h, v)
		}
	}
	var newHead ListNode
	point := &newHead
	for len(h) != 0 {
		res := (heap.Pop(&h)).(*ListNode)
		point.Next = res
		point = point.Next
		if res.Next != nil {
			heap.Push(&h, res.Next)
		}
		point.Next = nil
	}
	return newHead.Next

}

type ListNodeHeap []*ListNode

func (h ListNodeHeap) Len() int { return len(h) }

func (h ListNodeHeap) Less(i, j int) bool {
	return h[i].Val < h[j].Val
}

func (h ListNodeHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *ListNodeHeap) Push(x interface{}) {
	*h = append(*h, x.(*ListNode))
}

func (h *ListNodeHeap) Pop() interface{} {
	res := (*h)[len(*h)-1]
	*h = (*h)[:len(*h)-1]
	return res
}
