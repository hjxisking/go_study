package main

import (
    "fmt"
    "myleetcode/堆、栈与队列-前K个高频元素/normal"
)

func main() {
    //nums := []int{1,1,1,2,2,3}
    //k := 2

    //nums := []int{1}
    //k := 1

    nums := []int{4,1,-1,2,-1,2,3}
    k := 2
    fmt.Println(normal.TopKFrequent(nums, k))
}
