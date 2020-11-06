package main

import (
    "fmt"
    "myleetcode/数组与字符串-盛最多水的容器/normal"
)

func main() {
    //height := []int{1,8,6,2,5,4,8,3,7}
    //height := []int{1,1}
    height := []int{4,3,2,1,4}
    fmt.Println(normal.MaxArea(height))
}
