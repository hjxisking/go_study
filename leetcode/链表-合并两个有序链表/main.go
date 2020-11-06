package main

import (
    "fmt"
    "myleetcode/链表-合并两个有序链表/normal"
)

func main() {
    nums1 := []int{1,2,4}
    nums2 := []int{1,3,4}

    //nums1 := []int{1,2,4}
    //nums2 := []int{}

    l1 := makeChain(nums1)
    l2 := makeChain(nums2)

    print(normal.MergeTwoLists(l1, l2))
}

func makeChain(nums []int) *normal.ListNode {
    if len(nums) == 0 {
        return nil
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
