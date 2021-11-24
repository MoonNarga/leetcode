package solution

type MapSum struct {
	pair map[string]int
}

func Constructor_677() MapSum {
	p := MapSum{}
	p.pair = make(map[string]int)
	return p
}

func (this *MapSum) Insert(key string, val int) {
	this.pair[key] = val
}

func (this *MapSum) Sum(prefix string) int {
	sum := 0
	for i, v := range this.pair {
		if len(prefix) > len(i) {
			continue
		} else {
			tempSlice := i[:len(prefix)]
			if tempSlice == prefix {
				sum += v
			}
		}
	}
	return sum
}
