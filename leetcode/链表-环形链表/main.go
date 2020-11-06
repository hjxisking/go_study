package main

import (
    "fmt"
    "myleetcode/链表-环形链表/normal"
)

func main() {
    nums := []int{3,2,0,-4}
    head := makeChain(nums)
    pos := 1

    var circle *normal.ListNode
    node := head
    for i := 0; node.Next != nil; i ++ {
        if i == pos {
            circle = node.Next
        }

        node = node.Next
    }

    if circle != nil {
        node.Next = circle
    }

    fmt.Println(normal.HasCycle(head))
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
