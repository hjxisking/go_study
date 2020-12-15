package main

import (
    "fmt"
    "myleetcode/数组-旋转数组/normal"
)

func main() {
    //nums := []int{1,2,3,4,5,6,7}
    //k := 3

    //nums := []int{1,2,3,4}
    //k := 2

    nums := []int{1,2,3,4,5,6}
    k := 3

    //nums := []int{1,2}
    //k := 1
    normal.Rotate(nums, k)
    fmt.Println(nums)
}
