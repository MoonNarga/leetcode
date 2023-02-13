package solution

func balancedString(s string) int {
	avg := len(s) / 4
	str := []rune(s)
	cnt := map[rune]int{}
	maxCnt := 0
	for _, v := range str {
		cnt[v]++
		if cnt[v] > maxCnt {
			maxCnt = cnt[v]
		}
	}
	if maxCnt == avg {
		return 0
	}
	res := len(s)
	for i, j := 0, 0; j < len(str); j++ {
		cnt[str[j]]--
		for ; i <= j && cnt['Q'] <= avg && cnt['W'] <= avg && cnt['E'] <= avg && cnt['R'] <= avg; i++ {
			res = Min(res, j - i + 1)
			cnt[str[i]]++
		} 
	}
	return res
}
