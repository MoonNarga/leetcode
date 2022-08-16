package solution

type OrderedStream struct {
	order []string
	ptr   int
}

func Constructor(n int) OrderedStream {
	newOrderedStream := OrderedStream{order: make([]string, n), ptr: 0}
	return newOrderedStream
}

func (this *OrderedStream) Insert(idKey int, value string) []string {
	this.order[idKey-1] = value
	if this.ptr == idKey-1 {
		stream := make([]string, 0)
		for this.ptr < len(this.order) && this.order[this.ptr] != "" {
			stream = append(stream, this.order[this.ptr])
			this.ptr++
		}
		return stream
	}
	return []string{}
}

/**
 * Your OrderedStream object will be instantiated and called as such:
 * obj := Constructor(n);
 * param_1 := obj.Insert(idKey,value);
 */
