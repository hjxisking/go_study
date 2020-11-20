package main

import (
    "fmt"
    "myleetcode/binaryTree"
    "myleetcode/排序与搜索-二叉树中的最大路径和/normal"
)

const NIL = -999999999

func main() {
    //nums := []int{-10,9,20,-1,-1,15,7}
    //nums := []int{2,NIL,3}
    nums := []int{9,6,-3,NIL,NIL,-6,2,NIL,NIL,2,NIL,-6,-6,-6}
    root := makeTree(nums)
    fmt.Println(normal.MaxPathSum(root))
}

func makeTree(nums []int) *binaryTree.TreeNode {
    nodes := make([]*binaryTree.TreeNode, len(nums))
    for i, n := range nums {
        if n == NIL {
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
    root := nodes[0]
    sons := []*binaryTree.TreeNode{root}
    i := 1
    for len(sons) != 0 {
        tmp := []*binaryTree.TreeNode{}
        for _, node := range sons {
            if i < n && nodes[i] != nil {
                node.Left = nodes[i]
                tmp = append(tmp, nodes[i])
            }
            i ++
            if i < n && nodes[i] != nil {
                node.Right = nodes[i]
                tmp = append(tmp, nodes[i])
            }
            i ++
        }
        sons = tmp
    }

    return root
}