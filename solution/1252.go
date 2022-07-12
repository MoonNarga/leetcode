package solution

func oddCells(m int, n int, indices [][]int) int {
	countRow := make([]int, m)
	countCol := make([]int, n)
	for _, pair := range indices {
		x, y := pair[0], pair[1]
		countRow[x]++
		countCol[y]++
	}
	count := 0
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if (countRow[i]+countCol[j])%2 != 0 {
				count++
			}
		}
	}
	return count
}
