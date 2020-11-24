package main

import (
    "fmt"
    "myleetcode/回溯算法-全排列/backtrack"
    "myleetcode/回溯算法-全排列/normal"
)

func main() {
    nums := []int{1,2,3,4}
    fmt.Println(normal.Permute(nums))
    fmt.Println(backtrack.Permute(nums))
}
