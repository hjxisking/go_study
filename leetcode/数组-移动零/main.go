package main

import (
    "fmt"
    "myleetcode/数组-移动零/normal"
)

func main() {
    nums := []int{0,1,0,3,12}
    //nums := []int{1}

    normal.MoveZeroes(nums)
    fmt.Println(nums)
}
