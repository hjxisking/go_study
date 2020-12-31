package normal

// 无参考价值，标准解法应该使用【离散化树状数组】或者【归并排序】
func CountSmaller(nums []int) []int {
    ans := make([]int, len(nums))
    for i := 0; i < len(nums); i ++ {
        for j := i + 1; j < len(nums); j ++ {
            if nums[i] > nums[j] {
                ans[i] ++
            }
        }
    }
    return ans
}