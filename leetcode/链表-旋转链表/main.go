package main

import (
    "fmt"
    "myleetcode/链表-旋转链表/normal"
)

func main() {
    nums := []int{1,2}
    print(normal.RotateRight(makeChain(nums), 2))
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
        fmt.Print(current.Val, "->")
        next := current.Next
        if next == nil {
            break
        }
        current = next
    }
    fmt.Println("nil")
}
