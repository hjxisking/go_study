package normal

// 超时
func LengthOfLIS(nums []int) int {
    if len(nums) == 1 {
        return 1
    }

    ans := 0
    for i := 0; i < len(nums); i ++ {
        length := recursion(nums[i+1:], nums[i]) + 1
        if ans < length {
            ans = length
        }
    }

    return ans
}

func recursion(nums []int, n int) int {
    if len(nums) == 1 {
        if nums[0] > n {
            return 1
        } else {
            return 0
        }
    }
    ans := 0
    for i := 0; i < len(nums); i ++ {
        if nums[i] > n {
            tmp := recursion(nums[i+1:], nums[i]) + 1
            if tmp > ans {
                ans = tmp
            }
        }
    }
    return ans
}