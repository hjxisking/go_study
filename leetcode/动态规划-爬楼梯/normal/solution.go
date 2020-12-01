package normal

// f(n) = f(n-1) + f(n-2)
// 第N格楼梯可以由第N-1格直接走一步，或者第N-2格走两步得到（排除N-2格走两个一格的原因，是这种走法等于先走到N-1再走一格，已经包含在N-1格的走法中了）
// 所以第N格的走法其实就是第N-1的走法数量加上N-2的
func ClimbStairs(n int) int {
    ans := 0
    if n == 0 {
        return 0
    }
    if n == 1 {
        return 1
    }
    if n == 2 {
        return 2
    }

    prev := [2]int{1,2}
    for i := 3; i <= n; i ++ {
        ans = prev[0] + prev[1]
        prev[0] = prev[1]
        prev[1] = ans
    }

    return ans
}