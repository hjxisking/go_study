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

func OddEvenList(head *ListNode) *ListNode {
    if head == nil || head.Next == nil {
        return head
    }

    odd := head
    even, evenHead := head.Next, head.Next

    for even != nil && even.Next != nil {
        prevOdd := odd
        odd = even.Next
        prevOdd.Next = odd

        prevEven := even
        even = odd.Next
        prevEven.Next = even
    }

    odd.Next = evenHead
    return head
}