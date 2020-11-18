package main

import (
    "fmt"
    "myleetcode/排序与搜索-数组中的第K个最大元素/heapMax"
    "myleetcode/排序与搜索-数组中的第K个最大元素/heapMin"
    "myleetcode/排序与搜索-数组中的第K个最大元素/normal"
)

func main() {
    nums1 := []int{3,2,3,1,2,4,5,5,6,7,7,8,2,3,1,1,1,10,11,5,6,2,4,7,8,5,6}
    nums2 := []int{3,2,3,1,2,4,5,5,6,7,7,8,2,3,1,1,1,10,11,5,6,2,4,7,8,5,6}
    nums3 := []int{3,2,3,1,2,4,5,5,6,7,7,8,2,3,1,1,1,10,11,5,6,2,4,7,8,5,6}
    fmt.Println(normal.FindKthLargest(nums1, 20))
    fmt.Println(heapMax.FindKthLargest(nums2, 20))
    fmt.Println(heapMin.FindKthLargest(nums3, 20))
}
