package loop

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
    // 中序遍历，迭代的方式
    curr := root
    ans := -1
    stack := &stack{}

    for {
        for curr != nil {
            stack.push(curr)
            curr = curr.Left
        }
        if stack.len() == 0 {
            break
        } else {
            curr = stack.pop()
            k --
            if k == 0 {
                return curr.Val
            }
        }
        curr = curr.Right
    }
    return ans
}

type stack struct {
    nodes []*sortTree.TreeNode
}

func (s *stack) push(node *sortTree.TreeNode) {
    s.nodes = append(s.nodes, node)
}

func (s *stack) pop() *sortTree.TreeNode {
    n := s.len()
    if n == 0 {
        return nil
    }

    p := s.nodes[n-1]
    s.nodes = s.nodes[:n-1]

    return p
}

func (s *stack) len() int {
    return len(s.nodes)
}
