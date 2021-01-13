package main

import (
    "fmt"
    "myleetcode/动态规划-打家劫舍/loop"
    "myleetcode/动态规划-打家劫舍/normal"
)

func main() {
    //nums := []int{1,2,3,1}
    //nums := []int{2,7,9,3,1}
    nums := []int{2,1,1,2}
    fmt.Println(normal.Rob(nums))
    fmt.Println(loop.Rob(nums))
}
