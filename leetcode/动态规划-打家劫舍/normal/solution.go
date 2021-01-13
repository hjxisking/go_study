package normal

// 超时
func Rob(nums []int) int {
    switch len(nums) {
    case 0:
        return 0
    case 1:
        return nums[0]
    case 2:
        return max(nums[0], nums[1])
    }

    return max(nums[0] + Rob(nums[2:]), nums[1] + Rob(nums[3:]))
}

func max(x, y int) int {
    if x > y {
        return x
    }
    return y
}