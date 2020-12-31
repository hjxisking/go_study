package main

import (
    "fmt"
    "myleetcode/排序与检索-寻找重复数/normal"
)

func main() {
    //nums := []int{1,3,4,2,2}
    //nums := []int{1,4,3,3,3,5,3}
    nums := []int{2,2,2,2,2,2}
    fmt.Println(normal.FindDuplicate(nums))
}
