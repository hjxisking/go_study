package main

import (
    "fmt"
    "myleetcode/数组-两个数组的交集II/hash"
    "myleetcode/数组-两个数组的交集II/normal"
)

func main() {
    //nums1 := []int{1,2,2,1}
    //nums2 := []int{2,2}

    nums1 := []int{4,9,5}
    nums2 := []int{9,4,9,8,4}

    fmt.Println(normal.Intersect(nums1, nums2))
    fmt.Println(hash.Intersect(nums1, nums2))
}
