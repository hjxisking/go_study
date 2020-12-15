package normal

import "math"

func IncreasingTriplet(nums []int) bool {
    small, mid := math.MaxInt64, math.MaxInt64
    for i := 0; i < len(nums); i ++ {
        if mid > small && nums[i] > mid {
            return true
        }

        // 即使当前的mid比small大，替换后small虽然更小了，但是排序却在mid之后了，small可以被赋值的原因是：代表曾经有过一个比mid小，但是比small大的值在mid之前，如果后面遇到了更大值，那么也是成立的
        // small重新赋值后，下一个循环的nums[i]有2种可能
        // 1：小于small，那么就继续
        // 2：介于small和mid之间，那么形成了新的2连，且总体值更小
        // 3：大于mid，位置：mid，small（刚被换到mid后面去），max，也是成立的，因为small在mid之前有一个比现在small大且小于mid的值，现在又来个大于mid，可以推断 small(old), mid, small, max 也是成立的
        if nums[i] <= small {
            small = nums[i]
        }

        if nums[i] > small && nums[i] < mid {
            mid = nums[i]
        }
    }

    return false
}