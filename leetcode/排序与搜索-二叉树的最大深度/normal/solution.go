package normal

import "myleetcode/binaryTree"

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func MaxDepth(root *binaryTree.TreeNode) int {
    depth := 0
    if root == nil {
        return depth
    }
    queue := []*binaryTree.TreeNode{root}

    for len(queue) != 0 {
        depth ++
        tmp := make([]*binaryTree.TreeNode, 0)
        for _, node := range queue {
            if node.Left != nil {
                tmp = append(tmp, node.Left)
            }
            if node.Right != nil {
                tmp = append(tmp, node.Right)
            }
        }
        queue = tmp
    }

    return depth
}