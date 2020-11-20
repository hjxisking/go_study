package normal

import "myleetcode/binaryTree"

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val   int
 *     Left  *TreeNode
 *     Right *TreeNode
 * }
 */

func LowestCommonAncestor(root, p, q *binaryTree.TreeNode) *binaryTree.TreeNode {
    var min, max *binaryTree.TreeNode
    if p.Val < q.Val {
        min = p
        max = q
    } else {
        min = q
        max = p
    }

    curr := root
    for {
        if curr == nil {
            return nil
        }

        if curr.Val >= min.Val && curr.Val <= max.Val {
            break
        }

        if curr.Val < min.Val {
            curr = curr.Right
        }
        if curr.Val > max.Val {
            curr = curr.Left
        }
    }

    return curr
}