package doublePoints

func ProductExceptSelf(nums []int) []int {
    n := len(nums)
    i, j := 0, n - 1
    left, right := 1, 1
    ans := make([]int, n)
    for idx, _ := range ans {
        ans[idx] = 1
    }

    for i < n {
        ans[i] *= left
        ans[j] *= right
        left *= nums[i]
        right *= nums[j]

        i ++
        j --
    }

    return ans
}