package normal

import (
    "fmt"
    "math"
    "sort"
    "syscall"
)

func ThreeSumClosest(nums []int, target int) int {
    ans := [3]int{syscall.INFINITE/3, syscall.INFINITE/3, syscall.INFINITE/3}
    n := len(nums)
    sort.Sort(sort.IntSlice(nums))

    for first := 0; first < n - 2; first ++ {
        if first > 0 && nums[first] == nums[first-1] {
            continue
        }

        second := first + 1
        third := n - 1

        for third > second {
            if math.Abs(float64(nums[first] + nums[second] + nums[third] - target)) <= math.Abs(float64(ans[0] + ans[1] + ans[2] - target)) {
                ans[0] = nums[first]
                ans[1] = nums[second]
                ans[2] = nums[third]
            }

            if nums[first] + nums[second] + nums[third] > target {
                third --
            } else {
                second ++
            }
        }
    }
fmt.Println(ans)
    return ans[0] + ans[1] + ans[2]
}