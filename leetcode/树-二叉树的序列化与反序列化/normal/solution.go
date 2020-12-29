package normal

import (
    "strconv"
    "strings"
)

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

type TreeNode struct {
    Val int
    Left *TreeNode
    Right *TreeNode
}

const NIL = "nil"

type queue struct {
    data []*TreeNode
}

func (q *queue) push(node *TreeNode) {
    q.data = append(q.data, node)
}

func (q *queue) pop() *TreeNode {
    res := q.data[0]
    q.data = q.data[1:]
    return res
}

func (q *queue) popAll() []*TreeNode {
    res := q.data
    q.data = []*TreeNode{}
    return res
}

func (q *queue) isEmpty() bool {
    return len(q.data) == 0
}

func newQueue() *queue {
    return &queue{}
}

func breadthFirstSearch(root *TreeNode) []string {
    res := make([]string, 0)
    q := newQueue()
    q.push(root)
    res = append(res, strconv.Itoa(root.Val))
    nilCount := 0

    // nilCount的作用是，需要判断需要插入的nil节点是否为最后一层所有叶子节点产生的，如果是的话，这些nil节点不需要记录到结果集中去
    appendVal := func (node *TreeNode) {
        if nilCount > 0 {
            for i := 0; i < nilCount; i ++ {
                res = append(res, NIL)
            }
            nilCount = 0
        }
        res = append(res, strconv.Itoa(node.Val))
    }

    for !q.isEmpty() {
        current := q.pop()
        left := current.Left
        right := current.Right

        if left != nil {
            q.push(left)
            appendVal(left)
        } else {
            nilCount ++
        }

        if right != nil {
            q.push(right)
            appendVal(right)
        } else {
            nilCount ++
        }
    }

    return res
}

func newNode(n string) *TreeNode {
    if n == NIL {
        return nil
    }

    nVal, _ := strconv.Atoi(n)
    return &TreeNode{
        Val:   nVal,
        Left:  nil,
        Right: nil,
    }
}

func makeTree(tree []string) *TreeNode {
    if len(tree) == 0 {
        return nil
    }

    rootVal, _ := strconv.Atoi(tree[0])
    root := &TreeNode{
        Val:   rootVal,
        Left:  nil,
        Right: nil,
    }
    q := newQueue()
    q.push(root)

    getChild := func (i int) *TreeNode {
        var child *TreeNode
        if i < len(tree) {
            child = newNode(tree[i])
            if child != nil {
                q.push(child)
            }
        }
        return child
    }

    for lv, i := 0, 1; !q.isEmpty(); lv ++ {    // 按二叉树的层为单位循环
        parents := q.popAll()   // 取出该层所有父亲节点
        for _, parent := range parents {    // 为每个父亲节点寻找对应的左右子节点
            parent.Left = getChild(i)
            i ++
            parent.Right = getChild(i)
            i ++
        }
    }
    return root
}

type Codec struct {

}

func Constructor() Codec {
    return Codec{}
}

// Serializes a tree to a single string.
func (this *Codec) Serialize(root *TreeNode) string {
    if root == nil {
        return ""
    }
    return strings.Join(breadthFirstSearch(root), ",")
}

// Deserializes your encoded data to tree.
func (this *Codec) Deserialize(data string) *TreeNode {
    if data == "" {
        return nil
    }
    tree := strings.Split(data, ",")
    return makeTree(tree)
}


/**
 * Your Codec object will be instantiated and called as such:
 * ser := Constructor();
 * deser := Constructor();
 * data := ser.serialize(root);
 * ans := deser.deserialize(data);
 */