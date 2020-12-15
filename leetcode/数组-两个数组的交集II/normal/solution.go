package normal

import "sort"

func Intersect(nums1 []int, nums2 []int) []int {
    ans := []int{}
    sort.Sort(sort.IntSlice(nums1))
    sort.Sort(sort.IntSlice(nums2))
    n1 := len(nums1)
    n2 := len(nums2)
    i, j := 0, 0
    for i < n1 && j < n2 {
        if nums1[i] < nums2[j] {
            i ++
        } else if nums1[i] > nums2[j] {
            j ++
        } else {
            ans = append(ans, nums1[i])
            i ++
            j ++
        }
    }

    return ans
}

func min(x, y int) int {
    if x < y {
        return x
    }
    return y
}