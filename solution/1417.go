package solution

func reformat(s string) string {
	str := []byte(s)
	length := len(str)
	coNum, coLe := 0, 0
	for _, v := range str {
		if v > '9' {
			coLe++
		} else {
			coNum++
		}
	}
	if coLe-coNum > 1 || coLe-coNum < -1 {
		return ""
	}
	if coLe >= coNum {
		for l, r := 0, 1; r < length; l, r = l+2, r+2 {
			for ; r < length && str[r] <= '9'; r += 2 {
			}
			if r >= length {
				break
			}
			for ; l < length && str[l] > '9'; l += 2 {
			}
			str[l], str[r] = str[r], str[l]
		}
	} else {
		for l, r := 1, 0; r < length; l, r = l+2, r+2 {
			for ; l < length && str[l] > '9'; l += 2 {
			}
			if l >= length {
				break
			}
			for ; r < length && str[r] <= '9'; r += 2 {
			}
			str[l], str[r] = str[r], str[l]
		}
	}
	return string(str)
}
