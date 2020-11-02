package main

import (
    "fmt"
    "math"
)

func main() {
    //nums1 := []int{}
    //nums2 := []int{2}
    //nums1 := []int{1,3}
    //nums2 := []int{2}
    nums1 := []int{1,3}
    nums2 := []int{2,4}

    fmt.Println(findMedianSortedArrays(nums1, nums2))
}

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
    len1 := len(nums1)
    len2 := len(nums2)

    medianPos := int(math.Ceil(float64(len1+len2)/2))
    if (len1+len2)%2 == 0 {
        medianPos ++
    }

    idx1 := 0
    idx2 := 0
    m1 := 0
    m2 := 0
    for i := 0; i < medianPos; i ++ {
        m1 = m2
        if idx1 >= len1 {
            m2 = nums2[idx2]
            idx2 ++
        } else if idx2 >= len2 {
            m2 = nums1[idx1]
            idx1 ++
        } else if idx1 < len1 && idx2 < len2 {
            if nums1[idx1] < nums2[idx2] {
                m2 = nums1[idx1]
                idx1 ++
            } else {
                m2 = nums2[idx2]
                idx2 ++
            }
        }
    }

    if (len1+len2)%2 == 1 {
        m1 = m2
    }

    return float64(m1+m2)/2
}