package main

import (
    "fmt"
    "myleetcode/排序与搜索-搜索旋转排序数组/normal"
)

func main() {
    nums := []int{4,5,6,7,0,1,2}
    //nums := []int{1,3}
    //nums := []int{5,1,2,3,4}
    fmt.Println(normal.Search(nums, 1))
}
