package main

// 200. Number of Islands
func NumIslands(grid [][]byte) int {
	ROWs := len(grid)
	COLs := len(grid[0])
	res := 0
	visited := make(map[[2]int]bool, 0)
	var dfs func(r int, c int)
	dfs = func(r int, c int) {
		if r < 0 || r == ROWs || c < 0 || c == COLs || visited[[2]int{r, c}] || grid[r][c] == '0' {
			return
		}
		visited[[2]int{r, c}] = true
		dfs(r+1, c)
		dfs(r-1, c)
		dfs(r, c+1)
		dfs(r, c-1)
	}

	for r := 0; r < ROWs; r++ {
		for c := 0; c < COLs; c++ {
			if !visited[[2]int{r, c}] && grid[r][c] == '1' {
				dfs(r, c)
				res++
			}
		}
	}
	return res
}
