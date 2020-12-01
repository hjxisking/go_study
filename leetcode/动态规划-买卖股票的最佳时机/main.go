package main

import (
    "fmt"
    "myleetcode/动态规划-买卖股票的最佳时机/normal"
)

func main() {
    //nums := []int{7,1,5,3,6,4}
    nums := []int{7,6,5,4,3,1}
    fmt.Println(normal.MaxProfit(nums))
}
