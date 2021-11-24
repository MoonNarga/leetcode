package solution

func originalDigits(s string) string {
	count := make([]int, 26)
	nums := make([]int, 10)
	res := make([]byte, 40000)
	for _, v := range s {
		count[int(v)-int('a')]++
	}
	for count[25] > 0 {
		nums[0]++
		count[25]--
		count[int('e')-int('a')]--
		count[int('r')-int('a')]--
		count[int('o')-int('a')]--
	}
	for count[int('w')-int('a')] > 0 {
		nums[2]++
		count[int('t')-int('a')]--
		count[int('w')-int('a')]--
		count[int('o')-int('a')]--
	}
	for count[int('u')-int('a')] > 0 {
		nums[4]++
		count[int('f')-int('a')]--
		count[int('o')-int('a')]--
		count[int('u')-int('a')]--
		count[int('r')-int('a')]--
	}
	for count[int('f')-int('a')] > 0 {
		nums[5]++
		count[int('f')-int('a')]--
		count[int('i')-int('a')]--
		count[int('v')-int('a')]--
		count[int('e')-int('a')]--
	}
	for count[int('r')-int('a')] > 0 {
		nums[3]++
		count[int('t')-int('a')]--
		count[int('h')-int('a')]--
		count[int('r')-int('a')]--
		count[int('e')-int('a')] -= 2
	}
	for count[int('x')-int('a')] > 0 {
		nums[6]++
		count[int('s')-int('a')]--
		count[int('i')-int('a')]--
		count[int('x')-int('a')]--
	}
	for count[int('s')-int('a')] > 0 {
		nums[7]++
		count[int('s')-int('a')]--
		count[int('v')-int('a')]--
		count[int('n')-int('a')]--
		count[int('e')-int('a')]-=2
	}
	for count[int('t')-int('a')] > 0 {
		nums[8]++
		count[int('e')-int('a')]--
		count[int('i')-int('a')]--
		count[int('g')-int('a')]--
		count[int('h')-int('a')]--
		count[int('t')-int('a')]--
	}
	if count[int('o')-int('a')] > 0 {
		nums[1] += count[int('o')-int('a')]
		count[int('e')-int('a')] -= count[int('o')-int('a')]
	}
	nums[9] += count[int('e')-int('a')]
	c := 0
	for i, v := range nums {
		for j := 0; j < v; j++ {
			res[c] = byte(int('0')+i)
			c++
		}
	}
	return string(res[:c])
}
