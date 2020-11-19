package recursion

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
    if root == nil {
        return 0
    }

    return max(MaxDepth(root.Left), MaxDepth(root.Right)) + 1
}

func max(x, y int) int {
    if x > y {
        return x
    } else {
        return y
    }
}
