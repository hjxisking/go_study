package main

import (
    "fmt"
    "myleetcode/binaryTree"
    "myleetcode/binaryTree/sortTree"
    "myleetcode/排序与搜索-二叉搜索树的最近公共祖先/normal"
)

func main() {
    nums := []int{6,2,8,0,4,7,9,3,5}
    var root *binaryTree.TreeNode
    for _, num := range nums {
        root = sortTree.Insert(root, num)
    }
    p := root.Left.Right.Left
    q := root.Left.Right.Right

    fmt.Println(normal.LowestCommonAncestor(root, p, q))
}
