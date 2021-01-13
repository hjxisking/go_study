package main

import (
    "fmt"
    "myleetcode/动态规划-零钱兑换/normal"
)

func main() {
    //amount := 11
    //coins := []int{1,2,5}

    //amount := 3
    //coins := []int{2}

    //amount := 1
    //coins := []int{2}

    amount := 27
    coins := []int{2,5,10,1}
    fmt.Println(normal.CoinChange(coins, amount))
}
