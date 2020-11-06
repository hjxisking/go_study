package normal

import "math"

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

func MergeKLists(lists []*ListNode) *ListNode {
    head := &ListNode{}
    prev := head

    for {
        var tmp *ListNode
        var tmpIdx int
        min := int(math.Pow10(4))   // 题目的限制条件，排序内容小于 10^4
        for i := 0; i < len(lists); i ++ {
            if lists[i] != nil && lists[i].Val <= min {
                tmp = lists[i]
                min = tmp.Val
                tmpIdx = i
            }
        }

        if tmp == nil {
            break
        }

        prev.Next = tmp
        prev = tmp
        lists[tmpIdx] = lists[tmpIdx].Next
    }

    return head.Next
}