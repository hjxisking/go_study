package normal

import "myleetcode/binaryTree"

/**
 * Definition for TreeNode.
 * type TreeNode struct {
 *     Val int
 *     Left *ListNode
 *     Right *ListNode
 * }
 */
func LowestCommonAncestor(root, p, q *binaryTree.TreeNode) *binaryTree.TreeNode {
    if root == nil {
        return nil
    }

    if root.Val == p.Val || root.Val == q.Val {
        return root
    }

    left := LowestCommonAncestor(root.Left, p, q)
    right := LowestCommonAncestor(root.Right, p, q)

    if left != nil && right != nil {
        return root
    }

    if left == nil {
        return right
    }

    return left

}