// 二叉排序树
package sortTree

import "myleetcode/binaryTree"

func Insert(root *binaryTree.TreeNode, v int) *binaryTree.TreeNode {
    if root == nil {
        root = binaryTree.NewNode(v)
        return root
    }

    curr := root
    for curr != nil{
        if v < curr.Val {
            if curr.Left == nil {
                curr.Left = binaryTree.NewNode(v)
                break
            } else {
                curr = curr.Left
            }
        } else if v > curr.Val {
            if curr.Right == nil {
                curr.Right = binaryTree.NewNode(v)
                break
            } else {
                curr = curr.Right
            }
        } else {
            break
        }
    }

    return root
}

func Remove(root *binaryTree.TreeNode, v int) *binaryTree.TreeNode {
    curr := root
    prev := root
    for curr != nil {
        if v < curr.Val {
            prev = curr
            curr = curr.Left
        } else if v > curr.Val {
            prev = curr
            curr = curr.Right
        } else {
            var next *binaryTree.TreeNode
            if curr.Left == nil && curr.Right == nil {      // 叶子节点
                next = nil
            } else if curr.Left != nil && curr.Right == nil {   // 只有左子树
                next = curr.Left
            } else if curr.Left == nil && curr.Right != nil {   // 只有右子树
                next = curr.Right
            } else {    // 左右子树都有，那么寻找左子树的最大值，或者右子树的最小值，这个两个值是最接近被删除的curr，这里选择左子树的最大值
                for {
                    curr2 := curr
                    if curr2.Right == nil {
                        next = curr2
                        break
                    } else {
                        curr2 = curr2.Right
                    }
                }
            }

            if prev.Left == curr {
                prev.Left = next
            } else {
                prev.Right = next
            }

            break
        }
    }

    return root
}

// 前序遍历（递归）
func PreOrderRecursion(root *binaryTree.TreeNode) []int {
    return binaryTree.PreOrderRecursion(root)
}

// 前序遍历（迭代）
func PreOrder(root *binaryTree.TreeNode) []int {
    return binaryTree.PreOrder(root)
}

// 中序遍历（递归）
func InOrderRecursion(root *binaryTree.TreeNode) []int {
    return binaryTree.InOrderRecursion(root)
}

// 中序遍历（迭代）
func InOrder(root *binaryTree.TreeNode) []int {
    return binaryTree.InOrder(root)
}

// 后序遍历（递归）
func PostOrderRecursion(root *binaryTree.TreeNode) []int {
    return binaryTree.PostOrderRecursion(root)
}

// 后序遍历（迭代）
func PostOrder(root *binaryTree.TreeNode) []int {
    return binaryTree.PostOrder(root)
}

// 广度优先遍历
func BreadthFirst(root *binaryTree.TreeNode) []int {
    return binaryTree.BreadthFirst(root)
}