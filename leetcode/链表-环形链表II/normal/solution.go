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

func DetectCycle(head *ListNode) *ListNode {
    current, slow, fast := head, head, head

    for {
        if fast == nil || fast.Next == nil {
            return nil
        }

        fast = fast.Next.Next
        slow = slow.Next

        if fast == slow {
            break
        }
    }

    for current != slow {
        current = current.Next
        slow = slow.Next
    }

    return current
}
