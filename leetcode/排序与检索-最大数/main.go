package main

import (
    "fmt"
    "myleetcode/排序与检索-最大数/normal"
)

func main() {
    //nums := []int{3,30,34,5,9} // 9534330
    //nums := []int{432,43243}  // 43243432
    //nums := []int{8308,8308,830}  // 83088308830
    //nums := []int{8908,8908,890}  // 89089088908
    nums := []int{0,0}

    fmt.Println(normal.LargestNumber(nums))
}
