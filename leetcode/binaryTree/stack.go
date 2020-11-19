package binaryTree

type Stack struct {
    nodes []*TreeNode
}

func (s *Stack) Len() int {
    return len(s.nodes)
}

func (s *Stack) IsEmpty() bool {
    return s.Len() == 0
}

func (s *Stack) Push(node *TreeNode) {
    s.nodes = append(s.nodes, node)
}

func (s *Stack) Pop() *TreeNode {
    n := len(s.nodes)
    if n == 0 {
        return nil
    }

    node := s.nodes[n-1]
    s.nodes = s.nodes[:n-1]

    return node
}


func NewStack() *Stack {
    return &Stack{}
}