package main

import (
    "fmt"
    "myleetcode/数组-递增的三元子序列/normal"
)

func main() {
    //nums := []int{1,2,3,4,5}
    //nums := []int{5,4,3,2,1}
    //nums := []int{5,1,2,3,4}
    //nums := []int{2,1,5,0,4,6}
    //nums := []int{1,1,1,1,1,1,1,1}
    //nums := []int{5,1,5,5,2,5,4}
    //nums := []int{1,0,0,0,0,2,0,3}
    //nums := []int{1,2,-10,-8,-7}
    //nums := []int{7,8,5,1,2,3}
    //nums := []int{7,8,5,1,9,3}
    nums := []int{7,0,5,1,9,3}
    fmt.Println(normal.IncreasingTriplet(nums))
}
