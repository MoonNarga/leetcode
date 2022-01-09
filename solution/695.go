package solution

func dfs_695(grid [][]int, x, y, area int) int {
	area++
	grid[x][y] = 2
	if x > 0 && grid[x-1][y] == 1 {
		area = dfs_695(grid, x-1, y, area)
	}
	if x < len(grid)-1 && grid[x+1][y] == 1 {
		area = dfs_695(grid, x+1, y, area)
	}
	if y > 0 && grid[x][y-1] == 1 {
		area = dfs_695(grid, x, y-1, area)
	}
	if y < len(grid[x])-1 && grid[x][y+1] == 1 {
		area = dfs_695(grid, x, y+1, area)
	}
	return area
}

func maxAreaOfIsland(grid [][]int) int {
	maxArea := 0
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if grid[i][j] == 1 {
				temp := dfs_695(grid, i, j, 0)
				if temp > maxArea {
					maxArea = temp
				}
			}
		}
	}
	return maxArea
}
