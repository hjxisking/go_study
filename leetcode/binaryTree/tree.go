package binaryTree

type TreeNode struct {
    Val int
    Left *TreeNode
    Right *TreeNode
}

func NewNode(v int) *TreeNode {
    return &TreeNode{
        Val:   v,
        Left:  nil,
        Right: nil,
    }
}

// 前序遍历（递归）
func PreOrderRecursion(root *TreeNode) []int {
    list := make([]int, 0)
    if root == nil {
        return list
    }

    list = append(list, root.Val)
    left := PreOrderRecursion(root.Left)
    list = append(list, left...)
    right := PreOrderRecursion(root.Right)
    list = append(list, right...)

    return list
}

// 前序遍历（迭代）
func PreOrder(root *TreeNode) []int {
    list := make([]int, 0)
    if root == nil {
        return list
    }

    curr := root
    stack := NewStack()

    for {
        for curr != nil {
            list = append(list, curr.Val)
            stack.Push(curr)
            curr = curr.Left
        }

        if stack.IsEmpty() {
            break
        }

        curr = stack.Pop()
        curr = curr.Right
    }

    return list
}

// 中序遍历（递归）
func InOrderRecursion(root *TreeNode) []int {
    list := make([]int, 0)
    if root == nil {
        return list
    }

    left := InOrderRecursion(root.Left)
    list = append(list, left...)
    list = append(list, root.Val)
    right := InOrderRecursion(root.Right)
    list = append(list, right...)

    return list
}

// 中序遍历（迭代）
func InOrder(root *TreeNode) []int {
    list := make([]int, 0)
    if root == nil {
        return list
    }

    stack := NewStack()
    curr := root

    for {
        for curr != nil {
            stack.Push(curr)
            curr = curr.Left
        }

        if stack.IsEmpty() {
            break
        } else {
            curr = stack.Pop()
            list = append(list, curr.Val)
            curr = curr.Right
        }
    }

    return list
}

// 后序遍历（递归）
func PostOrderRecursion(root *TreeNode) []int {
    list := make([]int, 0)
    if root == nil {
        return list
    }

    left := PostOrderRecursion(root.Left)
    list = append(list, left...)
    right := PostOrderRecursion(root.Right)
    list = append(list, right...)
    list = append(list, root.Val)

    return list
}

// 后序遍历（迭代）
func PostOrder(root *TreeNode) []int {
    list := make([]int, 0)
    if root == nil {
        return list
    }

    stack := NewStack()
    var right *TreeNode
    curr := root

    loop:
    for {
        for curr != nil {
            stack.Push(curr)
            curr = curr.Left
        }

        curr = stack.Pop()

        // 如果curr是左叶子节点，那么当curr向上一个节点后，必然不满足等于right，所以左叶子节点进list
        // 如果curr是右叶子节点，那么当curr向上一个节点后，如果和当前右节点依然是右子关系，该右节点进入list，那么继续向上，直到关系变成左子关系便停止
        for curr.Right == nil || curr.Right == right {
            list = append(list, curr.Val)
            right = curr
            if stack.IsEmpty() {
                break loop
            }
            curr = stack.Pop()
        }

        // 这里有2种可能
        // 1. curr.Right != nil 直接进来的，那么说明还有右子树，那么pop出来的节点要先回去
        // 2. 是curr.Right == right 一路回退至curr变成stack.Pop()出来的左子节点，那么这时候也要假设是有右子树的，所以将pop出来的还原回去
        stack.Push(curr)
        curr = curr.Right
    }

    return list
}

// 广度优先遍历
func BreadthFirst(root *TreeNode) []int {
    list := make([]int, 0)
    if root == nil {
        return list
    }

    queue := NewQueue()
    queue.Push(root)
    for !queue.IsEmpty() {
        curr := queue.Pop()
        list = append(list, curr.Val)
        if curr.Left != nil {
            queue.Push(curr.Left)
        }
        if curr.Right != nil {
            queue.Push(curr.Right)
        }
    }

    return list
}