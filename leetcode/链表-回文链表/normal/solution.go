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

func findMiddle(head *ListNode) *ListNode {
    fast, slow := head, head
    for fast.Next != nil && fast.Next.Next != nil {
        fast = fast.Next.Next
        slow = slow.Next
    }
    return slow
}

// 如果不想破坏原来的链表结构，需要重新创建节点
func reverse(head *ListNode) *ListNode {
    current := head
    if current == nil || current.Next == nil {
        return head
    }

    var prev *ListNode
    // c -> p, p = c, c = n
    for current != nil {
        next := current.Next
        current.Next = prev
        prev = current
        current = next
    }
    return prev
}

// 快慢指针找到中点，然后反转后半部分，如果和前半部分一致就是回文
func IsPalindrome(head *ListNode) bool {
    if head == nil {
        return true
    }
    current := head
    mid := findMiddle(head)
    secondHalf := reverse(mid.Next)
    for secondHalf != nil {
        if current.Val != secondHalf.Val {
            return false
        }
        current = current.Next
        secondHalf = secondHalf.Next
    }
    return true
}