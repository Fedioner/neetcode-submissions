func numIslands(grid [][]byte) int {

	var ans int
	sizei, sizej := len(grid), len(grid[0])

	var dfs func(int, int) 

	dfs = func(i, j int){
		if i < 0 || j < 0 || i >= sizei || j>= sizej || grid[i][j] == '0' {
			return 
		}

		grid[i][j] = '0'

		dfs(i+1, j)
		dfs(i-1, j)
		dfs(i, j+1)
		dfs(i, j-1)
 	} 

	for i := 0; i < sizei; i++ {
		for j:= 0; j < sizej; j++ {
			if grid[i][j] == '1' {
				ans++
				dfs(i, j)
			}
		} 
	}

	return ans
}

