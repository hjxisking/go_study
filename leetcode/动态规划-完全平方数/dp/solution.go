package dp

// dp[0]=0
// dp[i] = for 1...squares { min(dp[i], dp[i-square] + 1) }
func NumSquares(n int) int {
    squares := []int{}
    for i := 1; i * i <= n; i ++ {
        squares = append(squares, i * i)
    }

    dp := make([]int, n + 1)
    for i, _ := range dp {
        dp[i] = 1 << 63 - 1
    }
    dp[0] = 0

    for i := 1; i < len(dp); i ++ {
        for _, square := range squares {
            if i < square {
                break
            }

            dp[i] = min(dp[i], dp[i-square] + 1)
        }
    }
    return dp[len(dp)-1]
}

func min(x, y int) int {
    if x < y {
        return x
    }
    return y
}