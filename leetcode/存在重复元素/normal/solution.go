package normal

import "sort"

func ContainsDuplicate(nums []int) bool {
    sort.Sort(sort.IntSlice(nums))
    n := len(nums)
    for i := 1; i < n; i ++ {
        if nums[i] == nums[i-1] {
            return true
        }
    }

    return false
}
