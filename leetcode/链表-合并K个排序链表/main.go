package main

import (
    "fmt"
    "myleetcode/链表-合并K个排序链表/normal"
)

func main() {
    lists := []*normal.ListNode{makeChain([]int{1,4,5}),makeChain([]int{1,3,4}),makeChain([]int{2,6})}
    print(normal.MergeKLists(lists))
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
