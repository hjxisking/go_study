package main

import (
    "fmt"
    "myleetcode/链表-环形链表II/normal"
)

func main() {
    nums := []int{1,2,3,4,5}
    head := makeChain(nums)
    pos := 1

    var circle *normal.ListNode
    node := head
    for i := 0; node.Next != nil; i ++ {
        if i == pos {
            circle = node
        }

        node = node.Next
    }

    if circle != nil {
        node.Next = circle
    }

    fmt.Println(normal.DetectCycle(head).Val)
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
