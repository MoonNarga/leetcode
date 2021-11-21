package solution

var co = [...]uint32{0, 1, 2, 4, 8, 16, 32, 64, 128, 256, 512, 1024, 2048, 4096, 8192, 16384, 32768, 65536, 131072, 262144, 524288, 1048576, 2097152, 4194304, 8388608, 16777216, 33554432}

func maxint(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func maxProduct(words []string) int {
	l := make([]int, len(words))
	code := make([]uint32, len(words))
	for i, v := range words {
		l[i] = len(v)
		for _, j := range v {
			ind := j - 96
			code[i] = code[i] | co[ind]
		}
	}
	max := 0
	for i := 0; i < len(code); i++ {
		for j := i + 1; j < len(code); j++ {
			if code[i]&code[j] != 0 {
				continue
			} else {
				max = maxint(max, l[i]*l[j])
			}
		}
	}
	return max
}
