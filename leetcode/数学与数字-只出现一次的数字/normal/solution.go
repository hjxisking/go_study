package normal

import (
    "sort"
)

func SingleNumber(nums []int) int {
    n := len(nums)
    if n == 1 {
        return nums[0]
    }

    if n == 2 {
        return 99999
    }

    sort.Sort(sort.IntSlice(nums))

    if nums[0] != nums[1] {
        return nums[0]
    }

    for i := 1; i < n - 1; i ++ {
        if nums[i] != nums[i-1] && nums[i] != nums[i+1] {
            return nums[i]
        }
    }

    if nums[n-1] != nums[n-2] {
        return nums[n-1]
    }

    return 99999
}