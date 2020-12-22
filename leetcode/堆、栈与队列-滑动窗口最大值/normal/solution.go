package normal

func MaxSlidingWindow(nums []int, k int) []int {
    ans := []int{}
    max := getMax(nums, 0, k)
    ans = append(ans, max)
    for i := k; i < len(nums); i ++ {   // i 是区间的右边界
        if nums[i-k] == max {   // 窗口的右边滑动到下标k的时候，如果移除的下标 i-k 正好是最大值，那么重新计算最大值
            max = getMax(nums, i-k+1, k)
        } else {
            if nums[i] > max {
                max = nums[i]
            }
        }
        ans = append(ans, max)
    }
    return ans
}

// 从i位置开始寻找k个位置中最大的值
func getMax(nums []int, i, k int) int {
    max := -1 << 63
    for j := i; j < k + i && j < len(nums); j ++ {
        if nums[j] > max {
            max = nums[j]
        }
    }
    return max
}