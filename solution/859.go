package solution

func buddyStrings(s string, goal string) bool {
	if len(s) != len(goal) {
		return false
	}
	var count [26]int
	var c1, c2, c3, c4 byte
	var flag, same int
	for i := 0; i < len(s); i++ {
		if s[i] == goal[i] {
			count[int(s[i])-int('a')]++
			same++
			continue
		} else {
			if flag == 0 {
				c1 = s[i]
				c2 = goal[i]
				flag++
			} else if flag == 1 {
				c3 = s[i]
				c4 = goal[i]
				flag++
			} else {
				return false
			}
		}
	}
	if flag == 0 {
		for _, v := range count {
			if v >= 2 {
				return true
			}
		}
		return false
	}
	if flag == 2 && c1 == c4 && c2 == c3 {
		return true
	}
	return false
}
