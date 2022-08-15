package solution

type MyCircularQueue struct {
	qu                        []int
	head, tail, length, usage int
}

func Constructor_622(k int) MyCircularQueue {
	newQu := MyCircularQueue{}
	newQu.length = k
	newQu.usage = 0
	newQu.qu = make([]int, k)
	newQu.head, newQu.tail = 0, 0
	return newQu
}

func (this *MyCircularQueue) EnQueue(value int) bool {
	if this.IsFull() {
		return false
	}
	this.qu[this.tail] = value
	this.usage++
	this.tail = (this.tail + 1) % this.length
	return true
}

func (this *MyCircularQueue) DeQueue() bool {
	if this.IsEmpty() {
		return false
	}
	this.usage--
	this.head = (this.head + 1) % this.length
	return true
}

func (this *MyCircularQueue) Front() int {
	if this.IsEmpty() {
		return -1
	}
	return this.qu[this.head]
}

func (this *MyCircularQueue) Rear() int {
	if this.IsEmpty() {
		return -1
	}
	return this.qu[(this.tail-1+this.length)%this.length]
}

func (this *MyCircularQueue) IsEmpty() bool {
	if this.usage == 0 {
		return true
	}
	return false
}

func (this *MyCircularQueue) IsFull() bool {
	if this.usage == this.length {
		return true
	}
	return false
}

/**
 * Your MyCircularQueue object will be instantiated and called as such:
 * obj := Constructor(k);
 * param_1 := obj.EnQueue(value);
 * param_2 := obj.DeQueue();
 * param_3 := obj.Front();
 * param_4 := obj.Rear();
 * param_5 := obj.IsEmpty();
 * param_6 := obj.IsFull();
 */
