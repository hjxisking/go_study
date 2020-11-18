package main

import (
    "fmt"
    "myleetcode/数学与数字-只出现一次的数字/normal"
)

func main() {
    //nums := []int{4,1,2,1,2}
    //nums := []int{5,5,4,1,2,1,2}
    nums := []int{2,2,1}
    fmt.Println(normal.SingleNumber(nums))
}
