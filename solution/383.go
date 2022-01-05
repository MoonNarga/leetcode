package solution

func canConstruct(ransomNote string, magazine string) bool {
	ran, mag := make([]int, 26), make([]int, 26)
	for _, v := range ransomNote {
		ran[int(v-'a')]++
	}
	for _, v := range magazine {
		mag[int(v-'a')]++
	}
	for i := 0; i < 26; i++ {
		if ran[i] > mag[i] {
			return false
		}
	}
	return true
}
