package main

import (
    "fmt"
    "myleetcode/数组与字符串-除自身意外数组的乘积/doublePoints"
    "myleetcode/数组与字符串-除自身意外数组的乘积/normal"
)

func main() {
    nums := []int{1,2,3,4}
    fmt.Println(normal.ProductExceptSelf(nums))
    fmt.Println(doublePoints.ProductExceptSelf(nums))
}
