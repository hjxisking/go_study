package main

import (
    "fmt"
    "myleetcode/排序与检索-摆动排序II/normal"
)

func main() {
    //nums := []int{1,5,1,1,6,4}
    //nums := []int{1,3,2,2,3,1}  // [2 3 1 3 1 2]
    //nums := []int{4,5,5,6}  // [5 6 4 5]
    //nums := []int{4,5,5,5,5,6,6,6}  // [5 6 5 6 5 6 4 5]
    nums := []int{5,3,1,2,6,7,8,5,5}    // [5 8 5 7 3 6 2 5 1]
    normal.WiggleSort(nums)
    fmt.Println(nums)
}
