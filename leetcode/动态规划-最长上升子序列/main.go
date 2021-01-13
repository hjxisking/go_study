package main

import (
    "fmt"
    "myleetcode/动态规划-最长上升子序列/dp"
    "myleetcode/动态规划-最长上升子序列/normal"
)

func main() {
    //nums := []int{10,9,2,5,3,7,101,18}
    //nums := []int{0,1,0,3,2,3}
    nums := []int{7,7,7,7,7,7,7,7}
    fmt.Println(normal.LengthOfLIS(nums))
    fmt.Println(dp.LengthOfLIS(nums))
}
