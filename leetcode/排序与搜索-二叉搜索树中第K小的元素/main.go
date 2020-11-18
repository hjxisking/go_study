package main

import (
    "fmt"
    "myleetcode/排序与搜索-二叉搜索树中第K小的元素/loop"
    "myleetcode/排序与搜索-二叉搜索树中第K小的元素/normal"
    "myleetcode/排序与搜索-二叉搜索树中第K小的元素/sortTree"
)

func main() {
    nums := []int{5,3,6,2,4,1}
    var root *sortTree.TreeNode
    for _, num := range nums {
        root = sortTree.Insert(root, num)
    }

    fmt.Println(normal.KthSmallest(root, 3))
    fmt.Println(loop.KthSmallest(root, 3))
}
