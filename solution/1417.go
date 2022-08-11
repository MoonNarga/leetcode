package solution

func reformat(s string) string {
	str := []byte(s)
	length := len(str)
	res := make([]byte, length)
	coNum, coLe := 0, 0
	num, letter := 0, 0
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
	if coNum >= coLe {
		num, letter = 0, 1
	} else {
		num, letter = 1, 0
	}
	for _, v := range str {
		if v > '9' {
			res[letter] = v
			letter += 2
		} else {
			res[num] = v
			num += 2
		}
	}
	return string(res)
}
