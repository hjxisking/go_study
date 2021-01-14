package dfs

var dirs = [][]int{{-1, 0}, {1, 0}, {0, 1}, {0, -1}}

func NumIslands(grid [][]byte) int {
    ans := 0
    rows := len(grid)
    columns := len(grid[0])
    for i := 0; i < rows; i ++ {
        for j := 0; j < columns; j ++ {
            if grid[i][j] == '1' {
                dfs(grid, i, j, rows, columns)
                ans ++
            }
        }
    }
    return ans
}

// 以x，y为起点，将岛屿所有的1都变成0
func dfs(grid [][]byte, x, y int, rows, columns int) {
    grid[x][y] = '0'
    for _, dir := range dirs {
        newX := x + dir[0]
        newY := y + dir[1]
        if newX >= 0 && newX < rows && newY >= 0 && newY < columns && grid[newX][newY] == '1' {
            dfs(grid, newX, newY, rows, columns)
        }
    }
}