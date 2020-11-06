package main

import (
    "fmt"
    "myleetcode/链表-两数相加/normal"
)

func main() {
    nums1 := []int{2,4,3}
    nums2 := []int{5,6,9}
    l1 := makeChain(nums1)
    l2 := makeChain(nums2)

    print(normal.AddTwoNumbers(l1, l2))
}

func makeChain(nums []int) *normal.ListNode {
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
    return head
}

func print(node *normal.ListNode) {
    current := node
    for {
        fmt.Print(current.Val)
        next := current.Next
        if next == nil {
            break
        }
        current = next
    }
    fmt.Println()
}