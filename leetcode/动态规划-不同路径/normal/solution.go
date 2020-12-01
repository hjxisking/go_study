package normal

func UniquePaths(m int, n int) int {
    path := make([][]int, m)
    for i := 0; i < m; i ++ {
        path[i] = make([]int, n)
        path[i][n-1] = 1
    }

    for i := n-2; i >= 0; i -- {
        path[0][i] = 1
    }

    for i := n-2; i >= 0; i -- {
        for j := 1; j < m; j ++ {
            path[j][i] = path[j][i+1] + path[j-1][i]
        }
    }

    return path[m-1][0]
}