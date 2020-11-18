package normal2

type ListNode struct {
    Val int
    Next *ListNode
}

func SortList(head *ListNode) *ListNode {
    if head == nil || head.Next == nil {
        return head
    }

    step := 1
    for {
        var h1, t1, h2, t2, lastTail *ListNode
        count := 0
        current := head
        for {
            if current == nil {
                break
            }

            h1, t1 = current, current
            for i := 0; i < step; i ++ {
                if current != nil {
                    count ++
                    t1 = current
                    current = current.Next
                } else {
                    break
                }
            }
            t1.Next = nil

            if current == nil {
                h2, t2 = nil, nil
            } else {
                h2, t2 = current, current
                for j := 0; j < step; j ++ {
                    if current != nil {
                        count ++
                        t2 = current
                        current = current.Next
                    } else {
                        break
                    }
                }
                t2.Next = nil
            }

            newHead, newTail := mergeTwoLists(h1, h2)
            if lastTail == nil {
                head = newHead
                lastTail = newTail
            } else {
                lastTail.Next = newHead
                lastTail = newTail
            }
        }

        if count <= step {
            break
        }

        step *= 2
    }

    return head
}

func mergeTwoLists(l1 *ListNode, l2 *ListNode) (*ListNode, *ListNode) {
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

    tail := prev
    for tail.Next != nil {
        tail = tail.Next
    }
    return head.Next, tail
}