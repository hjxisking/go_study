package normal

import (
    "math"
    "myleetcode/binaryTree"
)

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func MaxPathSum(root *binaryTree.TreeNode) int {
    maxSum := math.MinInt32
    var maxGain func(*binaryTree.TreeNode) int
    maxGain = func (node *binaryTree.TreeNode) int {
        if node == nil {
            return 0
        }

        leftGain := max(maxGain(node.Left), 0)
        rightGain := max(maxGain(node.Right), 0)

        maxSum = max(maxSum, node.Val + leftGain + rightGain)

        return max(leftGain, rightGain) + node.Val
    }
    maxGain(root)

    return maxSum
}


func max(x, y int) int {
    if x > y {
        return x
    } else {
        return y
    }
}