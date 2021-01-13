package loop

// [2, 7, 9, 3, 1]
// F(n) = Max( F(n-2) + A(n), F(n-1) )
// F(0) = A(0) = 2, F(1) = Max( A(0), A(1) ) = 7
// F(2) = Max( F(0) + A(2), F(1) )
func Rob(nums []int) int {
    switch len(nums) {
    case 0:
        return 0
    case 1:
        return nums[0]
    case 2:
        return max(nums[0], nums[1])
    }

    ans := 0
    f1, f2 := nums[0], max(nums[0], nums[1])
    for i := 2; i < len(nums); i ++ {
        ans = max( f1 + nums[i], f2 )
        f1 = f2
        f2 = ans
    }

    return ans
}

func max(x, y int) int {
    if x > y {
        return x
    }
    return y
}