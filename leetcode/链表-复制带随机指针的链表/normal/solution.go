package normal

/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Next *Node
 *     Random *Node
 * }
 */

type Node struct {
    Val int
    Next *Node
    Random *Node
}

func CopyRandomList(head *Node) *Node {
    if head == nil {
        return nil
    }

    nodeMap := map[*Node]*Node{}        // key：老的节点，value：新的节点
    newHead := &Node{}

    current := head
    newCurrent := newHead

    for current != nil {
        node := &Node{
            Val:    current.Val,
            Next:   nil,
            Random: nil,
        }
        nodeMap[current] = node
        current = current.Next
        newCurrent.Next = node
        newCurrent = node
    }

    current = head
    newCurrent = newHead.Next
    for current != nil{
        random := current.Random
        if random != nil {
            newCurrent.Random = nodeMap[random]
        }
        current = current.Next
        newCurrent = newCurrent.Next
    }

    return newHead.Next
}