package main

import (
    "fmt"
    "myleetcode/堆、栈与队列-滑动窗口最大值/deque"
    "myleetcode/堆、栈与队列-滑动窗口最大值/normal"
)

func main() {
    nums := []int{1,3,-1,-3,5,3,6,7}
    k := 3

    //nums := []int{1}
    //k := 1
    fmt.Println(normal.MaxSlidingWindow(nums, k))
    fmt.Println(deque.MaxSlidingWindow(nums, k))
}
