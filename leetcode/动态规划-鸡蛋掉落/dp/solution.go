package dp

// N：楼层数，K：鸡蛋数，X：在N层范围内的第几层丢鸡蛋
// dp[N][K] = min(dp[N][K], max(dp[N-X][K], dp[X-1][K-1])) + 1
func SuperEggDrop(K int, N int) int {
    // 行：层数（8，9，10）是3层，不是楼数
    // 列：鸡蛋数
    dp := make([][]int, N+1)
    for i, _ := range dp {
        dp[i] = make([]int, K+1)
    }

    // 初始化，所有dp结果均为最大值，即楼高N，因为最笨的办法就是逐层尝试，直到第N楼
    for i := 0; i <= N; i ++ {
        for j := 0; j <= K; j ++ {
            dp[i][j] = N
        }
    }

    // 0 行，即第0楼，无论多少格鸡蛋都不用试，移动次数为0
    // 1 行，即第1楼，无论鸡蛋多少颗，都只需要1次
    // 特殊情况：0列，没有鸡蛋的时候仍然为0
    for i := 0; i <= K; i ++ {
        dp[0][i] = 0
        dp[1][i] = 1
    }
    dp[1][0] = 0

    // 0 列，即没有鸡蛋，那么都不用尝试
    // 1 列，即只有一颗鸡蛋，那么最坏情况就是要逐层移动
    // 特殊情况：0层的时候不需要尝试
    for i := 0; i <= N; i ++ {
        dp[i][0] = 0
        dp[i][1] = i
    }
    dp[0][1] = 0

    for i := 2; i <= N; i ++ {
        for j := 2; j <= K; j ++ {
            left := 1
            right := i
            for left < right {
                mid := left + (right - left + 1) / 2
                lineUp := dp[mid-1][j-1]
                lineDonw := dp[i-mid][j]
                if lineUp == lineDonw {
                    left = mid
                    break
                } else if lineUp > lineDonw {
                    right = mid - 1
                } else {
                    left = mid
                }
            }
            x := left
            dp[i][j] = min(dp[i][j], max(dp[i-x][j], dp[x-1][j-1]) + 1)
        }
    }

    return dp[N][K]
}

func max(x, y int) int {
    if x > y {
        return x
    } else {
        return y
    }
}

func min(x, y int) int {
    if x < y {
        return x
    } else {
        return y
    }
}