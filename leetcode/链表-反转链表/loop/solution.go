package loop

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

func ReverseList(head *ListNode) *ListNode {
    if head == nil {
        return nil
    }

    current := head
    next := head.Next
    current.Next = nil

    for {
        if next == nil {
            break
        } else {
            tmpNode := next.Next
            next.Next = current
            current = next
            next = tmpNode
        }
    }

    return current
}