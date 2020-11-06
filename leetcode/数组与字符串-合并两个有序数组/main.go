package main

import (
    "fmt"
    "myleetcode/数组与字符串-合并两个有序数组/normal"
)

func main() {
    //nums1 := []int{1,2,8,0,0,0}
    //nums2 := []int{2,5,6}

    nums1 := []int{0}
    nums2 := []int{1}

    normal.Merge(nums1, 0, nums2, 1)
    fmt.Println(nums1)
}
