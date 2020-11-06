package main

import (
    "fmt"
    "myleetcode/链表-相交链表/normal"
)

func main() {
    numsI := []int{8,5,4}
    numsA := []int{4,1}
    numsB := []int{5,0,1}
    headI, _ := makeChain(numsI)
    headA, tailA := makeChain(numsA)
    headB, tailB := makeChain(numsB)

    tailA.Next = headI
    tailB.Next = headI

    fmt.Println(normal.GetIntersectionNode(headA, headB).Val)
}

func makeChain(nums []int) (*normal.ListNode, *normal.ListNode) {
    if len(nums) == 0 {
        return nil, nil
    }
    head := &normal.ListNode{
        Val:  nums[0],
        Next: nil,
    }
    current := head
    for i := 1; i < len(nums); i ++ {
        node := &normal.ListNode{
            Val:  nums[i],
            Next: nil,
        }
        current.Next = node
        current = node
    }
    return head, current
}