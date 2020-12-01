package normal

// f(i) = max(f(i-1)+Nums[i], Nums[i])
// 每次计算完f(i)需要赋值给Nums[i]，代表截止到i位置的最大值
// 例子：[-2,1,-3,4,-1,2,1,-5,4]
// f(0) = -2
// f(1) = 1 => nums[1] = 1 不变，同时代表的含义是，连续的最大子序和是从i=1开始的，因为带上i=0只会更小
// f(2) = -2 => nums[2] = -2，由原来的-3变成了-2，意思就是最大值是从1开始，截至到2，最大值为-2
// f(3) = 4 => nums[3] = 4 不变，解释同f(1)
// f(4) = 3 => nums[4] = 3，解释同f(2)
// 以此类推在f(7)的时候达到最大值（记录最大值需要用一个额外变量来记录）
func MaxSubArray(nums []int) int {
    max := nums[0]
    for i := 1; i < len(nums); i ++ {
        if nums[i] + nums[i-1] > nums[i] {
            nums[i] = nums[i] + nums[i-1]
        } else {
            // 代表连续子序的左边界变更成i
        }

        if nums[i] > max {
            max = nums[i]
            // 当前i可以是截至到目前位置的连续子序的右边界
        }
    }

    return max
}