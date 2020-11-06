package normal

import "sort"

func Merge(nums1 []int, m int, nums2 []int, n int)  {
    if n != 0 {
        for i := 0; i < m; i ++ {
            if nums1[i] > nums2[0] {
                tmp := nums2[0]
                nums2[0] = nums1[i]
                nums1[i] = tmp

                sort.Sort(sort.IntSlice(nums2))
            }
        }

        for j := 0; j < n; j ++ {
            nums1[m+j] = nums2[j]
        }
    }
}
