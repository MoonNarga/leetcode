package solution

func generateTheString(n int) string {
	res := make([]byte, n)
	if n%2 == 1 {
		for i := 0; i < n; i++ {
			res[i] = 'a'
		}
		return string(res)
	}
	for i := 0; i < n-1; i++ {
		res[i] = 'a'
	}
	res[n-1] = 'b'
	return string(res)
}
