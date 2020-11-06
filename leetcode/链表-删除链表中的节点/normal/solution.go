package normal

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

type ListNode struct {
    Val int
    Next *ListNode
}

func DeleteNode(node *ListNode) {
    prev := node
    for {
        if node.Next == nil {
            prev.Next = nil
            break
        }

        prev = node
        node.Val = node.Next.Val
        node = node.Next
    }
}