package main

import (
    "fmt"
    "myleetcode/动态规划-完全平方数/dp"
    "myleetcode/动态规划-完全平方数/greed"
    "myleetcode/动态规划-完全平方数/normal"
)

func main() {
    //n := 13
    //n := 12
    n := 62
    //n := 55
    fmt.Println(normal.NumSquares(n))
    fmt.Println(dp.NumSquares(n))
    fmt.Println(greed.NumSquares(n))
}
