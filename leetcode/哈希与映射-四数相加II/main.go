package main

import (
    "fmt"
    "myleetcode/哈希与映射-四数相加II/normal"
)

func main() {
    A := []int{ 1, 2}
    B := []int{-2,-1}
    C := []int{-1, 2}
    D := []int{ 0, 2}

    fmt.Println(normal.FourSumCount(A, B, C, D))
}
