package normal

func MissingNumber(nums []int) int {
    ans := 0
    n := len(nums)
    for i := 0; i < n; i ++ {
        ans += i - nums[i]
    }
    ans += n

    return ans
}