package solution

func isEdge(grid [][]int, row, col, w_grid, h_grid int) bool {
	c := grid[row][col]
	co := 0
	if row == 0 || grid[row-1][col] != c {
		co++
	}
	if col == 0 || grid[row][col-1] != c {
		co++
	}
	if row == h_grid-1 || grid[row+1][col] != c {
		co++
	}
	if col == w_grid-1 || grid[row][col+1] != c {
		co++
	}
	if co > 0 {
		return true
	}
	return false
}

func dfs_colorBorder(grid [][]int, visit_status, edge [][]bool, row, col int) {
	w_grid := len(grid[0])
	h_grid := len(grid)
	visit_status[row][col] = true
	c := grid[row][col]
	if row != 0 && grid[row-1][col] == c && !visit_status[row-1][col] {
		dfs_colorBorder(grid, visit_status, edge, row-1, col)
	}
	if col != 0 && grid[row][col-1] == c && !visit_status[row][col-1] {
		dfs_colorBorder(grid, visit_status, edge, row, col-1)
	}
	if row != h_grid-1 && grid[row+1][col] == c && !visit_status[row+1][col] {
		dfs_colorBorder(grid, visit_status, edge, row+1, col)
	}
	if col != w_grid-1 && grid[row][col+1] == c && !visit_status[row][col+1] {
		dfs_colorBorder(grid, visit_status, edge, row, col+1)
	}
	if isEdge(grid, row, col, w_grid, h_grid) {
		edge[row][col] = true
	}
}

func colorBorder(grid [][]int, row int, col int, color int) [][]int {
	visit_status := make([][]bool, len(grid))
	for i := 0; i < len(visit_status); i++ {
		visit_status[i] = make([]bool, len(grid[0]))
	}
	edge := make([][]bool, len(grid))
	for i := 0; i < len(edge); i++ {
		edge[i] = make([]bool, len(grid[0]))
	}
	dfs_colorBorder(grid, visit_status, edge, row, col)
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if edge[i][j] {
				grid[i][j] = color
			}
		}
	}
	return grid
}
