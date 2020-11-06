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

func HasCycle(head *ListNode) bool {
    if head == nil || head.Next == nil || head.Next.Next == nil {
        return false
    }

    slow, fast := head, head
    step := 2
    for {
        for i := 0; i < step; i ++ {
            fast = fast.Next
            if fast == nil {
                return false
            }

            if fast == slow {
                return true
            }
        }

        slow = slow.Next

        if fast == slow {
            return true
        }
    }
}