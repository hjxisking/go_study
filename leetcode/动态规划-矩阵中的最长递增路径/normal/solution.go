package normal

func LongestIncreasingPath(matrix [][]int) int {
    ans := 0
    rows := len(matrix)
    columns := len(matrix[0])
    cache := make([][]int, rows)
    for i := 0; i < len(matrix); i ++ {
        cache[i] = make([]int, columns)
    }

    for i := 0; i < len(matrix); i ++ {
        for j := 0; j < len(matrix[0]); j ++ {
            tmp := dfs(matrix, i, j, rows, columns, cache)
            if tmp > ans {
                ans = tmp
            }
        }
    }
    return ans
}

func dfs(matrix [][]int, i,j int, lx, ly int, cache [][]int) int {
    if cache[i][j] > 0 {
        return cache[i][j]
    }

    max := 0
    dirs := [][]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}
    for _, d := range dirs {
        row, column := i + d[0], j + d[1]
        if  row >= 0 && row < lx && column >= 0 && column < ly && matrix[row][column] > matrix[i][j] {
            res := dfs(matrix, row, column, lx, ly, cache)
            if res > max {
                max = res
            }
        }
    }
    max += 1
    cache[i][j] = max
    return max
}