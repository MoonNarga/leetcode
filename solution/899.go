package solution

func orderlyQueue(s string, k int) string {
	length := len(s)
	expand := make([]byte, length*2+1)
	sbyte := []byte(s)
	count := make([]int, 26)
	var minString, tempString string
	for i := 0; i < length; i++ {
		expand[i] = sbyte[i]
		count[int(sbyte[i])-int('a')]++
	}
	expand[length] = 0
	tempString = s
	minString = tempString
	if k == 1 {
		for i := 0; i < length-1; i++ {
			expand[length] = expand[0]
			expand = expand[1:]
			tempString = string(expand)
			if tempString < minString {
				minString = tempString
			}
		}
		return string([]byte(minString)[:length])
	}
	for i, j := 0, 0; i < 26; {
		if count[i] == 0 {
			i++
			continue
		}
		sbyte[j] = byte(int('a') + i)
		count[i]--
		j++
	}
	return string(sbyte)
}
