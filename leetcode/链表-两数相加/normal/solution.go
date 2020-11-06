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

func AddTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
    up := 0
    head := &ListNode{
        Val:  0,
        Next: nil,
    }

    current := head
    for {
       a := 0
       if l1 != nil {
           a = l1.Val
           l1 = l1.Next
       }

       b := 0
       if l2 != nil {
           b = l2.Val
           l2 = l2.Next
       }

       sum := a + b + up
       up = sum / 10
       current.Val = sum % 10

        next := &ListNode{
            Val:  0,
            Next: nil,
        }

        if l1 == nil && l2 == nil {
            if up > 0 {
                next.Val = up
                current.Next = next
            }
            break
        } else {
            current.Next = next
            current = next
        }
    }

    return head
}