package solution

func numSpecial(mat [][]int) int {
	row := make([]int, len(mat))
	col := make([]int, len(mat[0]))
	for i := 0; i < len(mat); i++ {
		for j := 0; j < len(mat[0]); j++ {
			if mat[i][j] == 1 {
				row[i]++
				col[j]++
			}
		}
	}
	res := 0
	for i := 0; i < len(mat); i++ {
		for j := 0; j < len(mat[0]); j++ {
			if mat[i][j] == 1 {
				if row[i] == 1 && col[j] == 1 {
					res++
				}
			}
		}
	}
	return res
}
