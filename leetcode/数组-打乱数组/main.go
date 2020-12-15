package main

import (
    "fmt"
    "myleetcode/数组-打乱数组/normal"
)

func main() {
    nums := []int{3,1,2,9,6,4,5,7}
    solution := normal.Constructor(nums)
    fmt.Println(solution.Shuffle())
    fmt.Println(solution.Reset())
    fmt.Println(solution.Shuffle())
}
