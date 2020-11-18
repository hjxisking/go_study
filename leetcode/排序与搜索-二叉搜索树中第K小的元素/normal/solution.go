package normal

import (
    "myleetcode/排序与搜索-二叉搜索树中第K小的元素/sortTree"
)

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func KthSmallest(root *sortTree.TreeNode, k int) int {
    list := ldr(root)
    return list[k-1]
}


// 递归
func ldr(node *sortTree.TreeNode) []int {
    list := []int{}

    if node.Left != nil {
        left := ldr(node.Left)
        list = append(list, left...)
    }

    list = append(list, node.Val)

    if node.Right != nil {
        right := ldr(node.Right)
        list = append(list, right...)
    }

    return list
}

