package bfs

var dirs = [][]int{{-1, 0}, {1, 0}, {0, 1}, {0, -1}}

func NumIslands(grid [][]byte) int {
    ans := 0
    rows := len(grid)
    columns := len(grid[0])
    for i := 0; i < rows; i ++ {
        for j := 0; j < columns; j ++ {
            if grid[i][j] == '1' {
                bfs(grid, i, j, rows, columns)
                ans ++
            }
        }
    }
    return ans
}

// 以x，y为起点进行广度优先遍历，将所有的1变成0
func bfs(grid [][]byte, x, y int, rows, columns int) {
    queue := [][]int{{x, y}}
    grid[x][y] = '0'
    for len(queue) > 0 {
        current := queue[0]
        queue = queue[1:]
        for _, dir := range dirs {
            newX := current[0] + dir[0]
            newY := current[1] + dir[1]
            if newX >= 0 && newX < rows && newY >= 0 && newY < columns && grid[newX][newY] == '1' {
                grid[newX][newY] = '0'
                queue = append(queue, []int{newX, newY})
            }
        }
    }
}