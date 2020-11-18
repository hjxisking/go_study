package normal

import "math"

func Reverse(x int) int {
    ans := 0
    nums := make([]int, 0, 10)
    max := int(math.Pow(2, 31)) - 1
    min := int(math.Pow(2, 31)) * -1

    for {
        nums = append(nums, x % 10)
        x = x / 10

        if x == 0 {
            break
        }
    }

    n := len(nums)

    if n > 10 {
        return 0
    }

    for i := 0; i < n; i ++ {
        if i == 0 && n == 10 {
            if math.Abs(float64(nums[i])) > 2 {
                return 0
            }
        }
        tmp := nums[i] * int(math.Pow10(n-1-i))
        if (tmp >= 0 && max - tmp - ans >= 0) || (tmp < 0 && min - tmp - ans <= 0) {
            ans += tmp
        } else {
            return 0
        }
    }

    return ans
}
