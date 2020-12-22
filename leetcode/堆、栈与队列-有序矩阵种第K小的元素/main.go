package main

import (
    "fmt"
    "myleetcode/堆、栈与队列-有序矩阵种第K小的元素/normal"
)

func main() {
    matrix := [][]int{{1,5,9},{10,11,13},{12,13,15}}
    k := 8
    fmt.Println(normal.KthSmallest(matrix, k))
}
