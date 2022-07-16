package main

import (
	"fmt"
	"container/list"
)

type MovingAverage struct {
	window *list.List
	cap   int
	sum   int
}

/** Initialize your data structure here. */
func Constructor(size int) MovingAverage {
	MA := MovingAverage{window: list.New(), cap: size, sum: 0}
	return MA
}

func (this *MovingAverage) Next(val int) float64 {
	this.window.PushBack(val)
	this.sum += val
	if this.cap < this.window.Len() {
		this.sum -= this.window.Front().Value.(int)
		this.window.Remove(this.window.Front())
	}
	return float64(this.sum) / float64(this.window.Len())
}

func main() {
	obj := Constructor(3)
	fmt.Println(obj.Next(1), obj.Next(10), obj.Next(3), obj.Next(5))
}
