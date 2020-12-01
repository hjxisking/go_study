package main

import (
    "fmt"
    "myleetcode/动态规划-最大子序和/normal"
)

func main() {
    nums := []int{-2,1,-3,4,-1,2,1,-5,4}
    fmt.Println(normal.MaxSubArray(nums))
}
