package main

import (
    "fmt"
    "myleetcode/搜索二维矩阵II/binarySearch"
    "myleetcode/搜索二维矩阵II/normal"
)

func main() {
    //matrix := [][]int{[]int{1,4,7,11,15},[]int{2,5,8,12,19},[]int{3,6,9,16,22},[]int{10,13,14,17,24},[]int{18,21,23,26,30}}
    //target := 17

    //matrix := [][]int{[]int{-1,3}}
    //target := 1

    //matrix := [][]int{[]int{-1},[]int{3}}
    //target := -1

    matrix := [][]int{[]int{4,9,9,11,12,15,17,22,23},[]int{8,10,11,14,14,17,20,23,26},[]int{9,13,16,21,23,26,30,35,36},[]int{9,14,18,21,26,29,32,35,39},[]int{12,18,19,23,27,33,34,37,39},[]int{15,20,24,26,32,34,36,39,42},[]int{16,24,28,28,33,37,37,43,44},[]int{18,28,28,29,35,42,44,49,52},[]int{23,32,34,34,39,46,51,51,56},[]int{27,35,35,40,45,47,51,55,58}}
    target := 30
    fmt.Println(normal.SearchMatrix(matrix, target))
    fmt.Println(binarySearch.SearchMatrix(matrix, target))
}
