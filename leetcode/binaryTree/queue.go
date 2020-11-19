package binaryTree

type Queue struct {
    nodes []*TreeNode
}

func NewQueue() *Queue {
    return &Queue{}
}

func (q *Queue) Push(node *TreeNode) {
    q.nodes = append(q.nodes, node)
}

func (q *Queue) Pop() *TreeNode {
    p := q.nodes[0]
    q.nodes = q.nodes[1:]

    return p
}

func (q *Queue) IsEmpty() bool {
    return len(q.nodes) == 0
}
