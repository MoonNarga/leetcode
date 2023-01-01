package solution

func repeatedCharacter(s string) byte {
	cnt := make([]int, 26)
	ss := []byte(s)
	for _, v := range ss {
		ord := v - 'a'
		cnt[ord]++
		if cnt[ord] == 2 {
			return v
		}
	}
	return 'a'
}
