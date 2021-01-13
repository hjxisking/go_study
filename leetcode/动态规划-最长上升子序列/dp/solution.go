package dp

// dp[i] = max(dp[j]) + 1，0 <= j < i and A[j] < A[i]
// dp[0] = 1
func LengthOfLIS(nums []int) int {
    ans := 0
    dp := make([]int, len(nums))
    dp[0] = 1

    for i := 1; i < len(nums); i ++ {
        max := 0
        for j := 0; j < i; j ++ {
            if nums[j] < nums[i] {
                if max < dp[j] {
                    max = dp[j]
                }
            }
        }
        dp[i] = max + 1
    }

    for i := 0; i < len(dp); i ++ {
        if ans < dp[i] {
            ans = dp[i]
        }
    }

    return ans
}