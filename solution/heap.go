package solution

type Heap[T any] struct {
	data []T
	cmp  func(a, b T) bool
}

func NewHeap[T any](cmp func(a, b T) bool) *Heap[T] {
	return &Heap[T]{
		data: make([]T, 0),
		cmp:  cmp,
	}
}

func (h *Heap[T]) swap(i, j int) {
	h.data[i], h.data[j] = h.data[j], h.data[i]
}

func (h *Heap[T]) swim(i int) {
	for h.cmp(h.data[(i-1)/2], h.data[i]) {
		h.swap((i-1)/2, i)
		i = (i - 1) / 2
	}
}

func (h *Heap[T]) sink(i int) {
	l, r := i*2+1, i*2+2
	largest := i
	if l < len(h.data) && h.cmp(h.data[i], h.data[l]) {
		largest = l
	}
	if r < len(h.data) && h.cmp(h.data[largest], h.data[r]) {
		largest = r
	}
	if largest != i {
		h.swap(i, largest)
		h.sink(largest)
	}
}

func (h *Heap[T]) Replace(v T) (T, bool) {
	if h.Size() == 0 {
		return v, false
	}
	h.data[0], v = v, h.data[0]
	h.sink(0)
	return v, true
}

func (h *Heap[T]) Size() int {
	return len(h.data)
}

func (h *Heap[T]) Peek() (T, bool) {
	if h.Size() == 0 {
		var v T
		return v, false
	}
	return h.data[0], true
}

func (h *Heap[T]) Push(v T) {
	h.data = append(h.data, v)
	h.swim(h.Size() - 1)
}

func (h *Heap[T]) Pop() (T, bool) {
	var v T
	if h.Size() == 0 {
		return v, false
	}
	v = h.data[0]
	h.swap(0, h.Size()-1)
	h.data = h.data[:h.Size()-1]
	h.sink(0)
	return v, true
}
