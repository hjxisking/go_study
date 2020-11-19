package main

import (
    "fmt"
    "myleetcode/binaryTree"
    "myleetcode/binaryTree/sortTree"
)

func main() {
    nums := []int{20,10,30,7,15,25,40,6,8,13,18,3}
    var root *binaryTree.TreeNode
    for _, n := range nums {
        root = sortTree.Insert(root, n)
    }
    fmt.Println("前序遍历")
    fmt.Println(sortTree.InOrderRecursion(root))
    fmt.Println(sortTree.InOrder(root))
    fmt.Println()
    fmt.Println("中序遍历")
    fmt.Println(sortTree.PreOrderRecursion(root))
    fmt.Println(sortTree.PreOrder(root))
    fmt.Println()
    fmt.Println("后序遍历")
    fmt.Println(sortTree.PostOrderRecursion(root))
    fmt.Println(sortTree.PostOrder(root))
    fmt.Println("层序遍历")
    fmt.Println(sortTree.BreadthFirst(root))
}