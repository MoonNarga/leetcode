package solution

import (
	"math/rand"
	"time"
)

type Solution struct {
	origin   []int
	outorder []int
}

func Constructor_384(nums []int) Solution {
	s := Solution{}
	for _, v := range nums {
		s.origin = append(s.origin, v)
		s.outorder = append(s.outorder, v)
	}
	return s
}

func (this *Solution) Reset() []int {
	return this.origin
}

func (this *Solution) Shuffle() []int {
	rand.Seed(time.Now().UnixNano())
	l := len(this.outorder)
	for i := 0; i < l-1; i++ {
		r := rand.Intn(l - i)
		this.outorder[r], this.outorder[l-1-i] = this.outorder[l-1-i], this.outorder[r]
	}
	return this.outorder
}

/**
 * Your Solution object will be instantiated and called as such:
 * obj := Constructor(nums);
 * param_1 := obj.Reset();
 * param_2 := obj.Shuffle();
 */
