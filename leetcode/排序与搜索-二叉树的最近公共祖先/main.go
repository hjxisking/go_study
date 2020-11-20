package main

import (
    "fmt"
    "math"
    "myleetcode/binaryTree"
    "myleetcode/排序与搜索-二叉树的最近公共祖先/normal"
)

const NIL = math.MinInt32

func main() {
    nums := []int{3,5,1,6,2,0,8,NIL,NIL,7,4}
    root := makeTree(nums)
    p := root.Left.Left
    q := root.Left.Right.Left

    fmt.Println(normal.LowestCommonAncestor(root, p, q))
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