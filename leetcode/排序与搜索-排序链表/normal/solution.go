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

func SortList(head *ListNode) *ListNode {
    // 只有一个节点，立刻返回
    if head == nil || head.Next == nil {
        return head
    }

    // 只有两个节点，进行排序后返回
    if head.Next.Next == nil {
        current := head
        next := head.Next
        if current.Val > next.Val {
            next.Next = current
            current.Next = nil
            return next
        } else {
            return current
        }
    }

    slow, fast, left, right := head, head, head, head
    for {
        if fast.Next != nil {
            fast = fast.Next
            if fast.Next != nil {
                fast = fast.Next
            } else {
                break
            }
        } else {
            break
        }

        slow = slow.Next
    }

    right = slow.Next
    slow.Next = nil

    left = SortList(left)
    right = SortList(right)

    return mergeTwoLists(left, right)
}

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
    head := &ListNode{}
    prev := head

    for l1 != nil && l2 != nil {
        if l1.Val <= l2.Val {
            prev.Next = l1
            l1 = l1.Next
        } else {
            prev.Next = l2
            l2 = l2.Next
        }

        prev = prev.Next
    }

    if l1 == nil {
        prev.Next = l2
    } else {
        prev.Next = l1
    }

    return head.Next
}
