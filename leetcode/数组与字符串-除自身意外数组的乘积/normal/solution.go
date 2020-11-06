package normal

func ProductExceptSelf(nums []int) []int {
    n := len(nums)
    prefix := make([]int, n)
    postfix := make([]int, n)
    ans := make([]int, n)

    prefix[0] = nums[0]
    for i := 1; i < n - 1; i ++ {
        prefix[i] = nums[i] * prefix[i-1]
    }

    postfix[n-1] = nums[n-1]
    for i := n - 2; i >= 0; i -- {
        postfix[i] = nums[i] * postfix[i+1]
    }

    ans[0] = postfix[1]
    ans[n-1] = prefix[n-2]
    for i := 1; i < n - 1; i ++ {
        ans[i] = prefix[i-1] * postfix[i+1]
    }
    return ans
}