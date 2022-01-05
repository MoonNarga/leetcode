package solution

func modifyString(s string) string {
	sbytes := []byte(s)
	if len(sbytes) == 1 {
		if sbytes[0] == '?' {
			sbytes[0] = 'a'
		}
	}
	if sbytes[0] == '?' {
		for i := 0; i < 26; i++ {
			if byte(i+int('a')) != sbytes[1] {
				sbytes[0] = byte(i + int('a'))
				break
			}
		}
	}
	for i := 1; i < len(sbytes)-1; i++ {
		if sbytes[i] == '?' {
			for j := 0; j < 26; j++ {
				if byte(j+int('a')) != sbytes[i-1] && byte(j+int('a')) != sbytes[i+1] {
					sbytes[i] = byte(j + int('a'))
					break
				}
			}
		}
	}
	if sbytes[len(s)-1] == '?' {
		for i := 0; i < 26; i++ {
			if byte(i+int('a')) != sbytes[len(s)-2] {
				sbytes[len(s)-1] = byte(i + int('a'))
				break
			}
		}
	}
	s = string(sbytes)
	return s
}
