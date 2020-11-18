package sortTree

type TreeNode struct {
    Val int
    Left *TreeNode
    Right *TreeNode
}

func newTreeNode(i int) *TreeNode {
    return &TreeNode{
        Val:   i,
        Left:  nil,
        Right: nil,
    }
}

func Insert(root *TreeNode, i int) *TreeNode {
    node := newTreeNode(i)

    if root == nil {
        root = node
        return root
    }

    curr := root
    for {
        if i < curr.Val {
            if curr.Left == nil {
                curr.Left = node
                break
            } else {
                curr = curr.Left
            }
        } else if i > curr.Val {
            if curr.Right == nil {
                curr.Right = node
                break
            } else {
                curr = curr.Right
            }
        } else {    // 说明节点存在
            return root
        }
    }

    return root
}

func Ldr(node *TreeNode) []int {
    list := []int{}

    if node.Left != nil {
        left := Ldr(node.Left)
        list = append(list, left...)
    }

    list = append(list, node.Val)

    if node.Right != nil {
        right := Ldr(node.Right)
        list = append(list, right...)
    }

    return list
}