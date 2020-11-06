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

func RotateRight(head *ListNode, k int) *ListNode {
    if head == nil || head.Next == nil {
        return head
    }

    tail := &ListNode{}
    tail.Next = head
    length := 1

    node := head
    for {
        k --
        if node.Next == nil {
            break
        }

        length ++
        if k <= 0 {
            tail = tail.Next
        }
        node = node.Next
    }

    if k % length == 0 {
        return head
    }

    if k > 0 {
        k = k % length
        tail = head
        for i := 0; i < length - k - 1; i ++ {
            tail = tail.Next
        }
    }

    node.Next = head
    head = tail.Next
    tail.Next = nil

    return head
}