package main

import (
    "fmt"
    "myleetcode/排序与检索-寻找峰值/binarySearch"
    "myleetcode/排序与检索-寻找峰值/normal"
)

func main() {
    nums := []int{1,2,1,3,5,6,4}
    //nums := []int{2,1}
    //nums := []int{1,2}
    fmt.Println(normal.FindPeakElement(nums))
    fmt.Println(binarySearch.FindPeakElement(nums))
}
