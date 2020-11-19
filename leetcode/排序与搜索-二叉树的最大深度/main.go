package main

import (
    "fmt"
    "myleetcode/binaryTree"
    "myleetcode/排序与搜索-二叉树的最大深度/normal"
    "myleetcode/排序与搜索-二叉树的最大深度/recursion"
)

func main() {
    nums := []int{3,9,20,-1,-1,15,7}
    root := makeTree(nums)
    fmt.Println(normal.MaxDepth(root))
    fmt.Println(recursion.MaxDepth(root))
}

func makeTree(nums []int) *binaryTree.TreeNode {
    nodes := make([]*binaryTree.TreeNode, len(nums))
    for i, n := range nums {
        if n == -1 {
            nodes[i] = nil
        } else {
            nodes[i] = &binaryTree.TreeNode{
                Val:   n,
                Left:  nil,
                Right: nil,
            }
        }
    }

    n := len(nodes)
    for i, _ := range nodes {
        if n > 2*i+1 {
            nodes[i].Left = nodes[2*i+1]
        }
        if n > 2*i+2 {
            nodes[i].Right = nodes[2*i+2]
        }
    }

    return nodes[0]
}
