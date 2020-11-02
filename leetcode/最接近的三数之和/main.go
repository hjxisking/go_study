package main

import (
    "fmt"
    "myleetcode/最接近的三数之和/normal"
)

func main() {
    //nums := []int{-1,2,1,-4}
    //target := 1

    // -1
    nums := []int{1,1,-1,-1,3}
    target := -1

    //nums := []int{0,2,1,-3}
    //target := 1

    // 2
    //nums := []int{-1,2,1,-4}
    //target := 1

    // 13
    //nums := []int{1,2,5,10,11}
    //target := 12

    fmt.Println(normal.ThreeSumClosest(nums, target))
}
