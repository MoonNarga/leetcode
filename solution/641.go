package solution

type MyCircularDeque struct {
	deque        []int
	front, rear  int
	lengeth, cap int
}

func Constructor_641(k int) MyCircularDeque {
	newDeque := MyCircularDeque{deque: make([]int, k), front: 0, rear: 0, lengeth: 0, cap: k}
	return newDeque
}

func (this *MyCircularDeque) InsertFront(value int) bool {
	if this.IsFull() {
		return false
	}
	if this.IsEmpty() {
		this.deque[this.front] = value
	} else {
		this.front = (this.front - 1 + this.cap) % this.cap
		this.deque[this.front] = value
	}
	this.lengeth++
	return true
}

func (this *MyCircularDeque) InsertLast(value int) bool {
	if this.IsFull() {
		return false
	}
	if this.IsEmpty() {
		this.deque[this.rear] = value
	} else {
		this.rear = (this.rear + 1) % this.cap
		this.deque[this.rear] = value
	}
	this.lengeth++
	return true
}

func (this *MyCircularDeque) DeleteFront() bool {
	if this.IsEmpty() {
		return false
	}
	if this.lengeth != 1 {
		this.front = (this.front + 1) % this.cap
	}
	this.lengeth--
	return true
}

func (this *MyCircularDeque) DeleteLast() bool {
	if this.IsEmpty() {
		return false
	}
	if this.lengeth != 1 {
		this.rear = (this.rear - 1 + this.cap) % this.cap
	}
	this.lengeth--
	return true
}

func (this *MyCircularDeque) GetFront() int {
	if this.IsEmpty() {
		return -1
	}
	return this.deque[this.front]
}

func (this *MyCircularDeque) GetRear() int {
	if this.IsEmpty() {
		return -1
	}
	return this.deque[this.rear]
}

func (this *MyCircularDeque) IsEmpty() bool {
	return this.lengeth == 0
}

func (this *MyCircularDeque) IsFull() bool {
	return this.lengeth == this.cap
}

/**
 * Your MyCircularDeque object will be instantiated and called as such:
 * obj := Constructor(k);
 * param_1 := obj.InsertFront(value);
 * param_2 := obj.InsertLast(value);
 * param_3 := obj.DeleteFront();
 * param_4 := obj.DeleteLast();
 * param_5 := obj.GetFront();
 * param_6 := obj.GetRear();
 * param_7 := obj.IsEmpty();
 * param_8 := obj.IsFull();
 */
