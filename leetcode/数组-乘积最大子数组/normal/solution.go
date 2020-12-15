package normal

import "math"

// Fmax(i) = Max( Fmax(i-1)*nums[i], Fmin(i-1)*nums[i], nums[i] )
// Fmin(i) = Min( Fmin(i-1)*nums[i], Fmin(i-1)*nums[i], nums[i] )
// 引入nums[i]的原因是，有可能前面的值是0
func MaxProduct(nums []int) int {
    Fmax := make([]int, len(nums))
    Fmin := make([]int, len(nums))

    Fmax[0] = nums[0]
    Fmin[0] = nums[0]
    for i := 1; i < len(nums); i ++ {
        Fmax[i] = max( Fmax[i-1]*nums[i], max(Fmin[i-1]*nums[i], nums[i]) )
        Fmin[i] = min( Fmax[i-1]*nums[i], min(Fmin[i-1]*nums[i], nums[i]) )
    }

    ans := math.MinInt64
    for i := 0; i < len(Fmax); i ++ {
        ans = max(Fmax[i], ans)
    }

    return ans
}

func max(x, y int) int {
    if x > y {
        return x
    } else {
        return y
    }
}

func min(x, y int) int {
    if x < y {
        return x
    } else {
        return y
    }
}

