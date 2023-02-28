package solution

func parseInt(str []rune) int {
	res := 0
	for _, v := range str {
		res *= 10
		res += int(v - '0')
	}
	return res
}

func calculate(s string) int {
	str := []rune(s)
	nums := []int{}
	symbols := []rune{}
	start := true
	for i := 0; i < len(str); {
		switch str[i] {
		case ' ':
			i++
			continue
		case '(':
			symbols = append(symbols, str[i])
			start = true
			i++
		case '-':
			if start {
				nums = append(nums, 0)
			}
			fallthrough
		case '+':
			symbols = append(symbols, str[i])
			i++
			start = false
			continue
		case ')':
			symbols = symbols[:len(symbols)-1]
			i++
			start = false
		default:
			start = false
			temp := i
			for i < len(str) && str[i] >= '0' && str[i] <= '9' {
				i++
			}
			nums = append(nums, parseInt(str[temp:i]))
		}
		if len(symbols) != 0 && symbols[len(symbols)-1] != '(' {
			if symbols[len(symbols)-1] == '-' {
				nums[len(nums)-2] -= nums[len(nums)-1]
			} else {
				nums[len(nums)-2] += nums[len(nums)-1]
			}
			nums = nums[:len(nums)-1]
			symbols = symbols[:len(symbols)-1]
		}
	}
	return nums[0]
}
