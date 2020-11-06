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

func GetIntersectionNode(headA, headB *ListNode) *ListNode {
    if headA == nil || headB == nil {
        return nil
    }

    nodeA := headA
    nodeB := headB

    lenA, lenB := 0, 0
    for nodeA != nil || nodeB != nil {
        if nodeA != nil {
            nodeA = nodeA.Next
            lenA ++
        }

        if nodeB != nil {
            nodeB = nodeB.Next
            lenB ++
        }
    }

    var first, second *ListNode
    distance := 0
    if lenA > lenB {
        first = headA
        second = headB
        distance = lenA - lenB
    } else {
        first = headB
        second = headA
        distance = lenB - lenA
    }

    for {
        if first == second || first == nil {
            break
        }

        if distance == 0 {
            second = second.Next
        } else {
            distance --
        }

        first = first.Next
    }

    return first
}