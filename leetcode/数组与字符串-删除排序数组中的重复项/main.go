package main

import (
    "fmt"
    "myleetcode/数组与字符串-删除排序数组中的重复项/normal"
)

func main() {
    nums := []int{0,0,1,1,1,1,2,2,3,3,4}
    fmt.Println(normal.RemoveDuplicates(nums), nums)
}
