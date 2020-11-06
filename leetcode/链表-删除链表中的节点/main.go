package main

import (
    "fmt"
    "myleetcode/链表-删除链表中的节点/normal"
)

func main() {
    nums := []int{4,5,1,9}
    head, delete := makeChain(nums, 2)
    print(head)
    normal.DeleteNode(delete)
    print(head)
}

func makeChain(nums []int, p int) (*normal.ListNode, *normal.ListNode) {
    if len(nums) == 0 {
        return nil, nil
    }
    var ret *normal.ListNode
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
        if p == i {
            ret = node
        }
    }
    return head, ret
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
