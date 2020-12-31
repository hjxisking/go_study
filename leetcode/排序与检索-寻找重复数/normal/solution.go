package normal

// 思路：二分查找
// 假设 nums = [1,4,3,3,3,5,3]，最大自然数n为6 （len=6+1）
// nums中
//  小于等于 6 的个数是 7
//  小于等于 5 的个数是 7
//  小于等于 4 的个数是 6
//  小于等于 3 的个数是 5
//  小于等于 2 的个数是 1
//  小于等于 1 的个数是 1
// 也就是，自然数范围 1 - 6 内，逐个检查nums内小于等于该自然数的count，如果没有重复数字，那么小于自然数 i 的个数一定也是小于等于 i，只有出现重复数字之后，count数量才会上升并大于 i
// 那么问题就可以演化成如何快速找到 2 ？
// 从规律中可以发现，自从 2 是重复的之后，计算3的count数量也因为 2 多了一个而count变大了，后续数字以此类推，那么只要找到第一个变大的数字即可
// 二分查找法：
//      找到中点（1 - 6）mid=3，如果自然数3的count数量小于等于3，那么继续二分右边，如果count数量大于3，那么继续二分左边（同时不排除当前mid正好就是答案，也就是第一个开始变大的点，所以要记录下来）
func FindDuplicate(nums []int) int {
    ans := -1
    n := len(nums)
    left, right := 0, n - 1
    for left <= right {
        // mid 是 1 至 n - 1 的连续自然数的中间数
        mid := (left + right) / 2
        count := 0
        for i := 0; i < n; i ++ {
            if nums[i] <= mid {
                count ++
            }
        }
        if count <= mid {
            left = mid + 1
        } else {
            right = mid - 1
            ans = mid
        }
    }
    return ans
}