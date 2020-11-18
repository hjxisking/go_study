package main

import (
    "fmt"
    //normal "myleetcode/排序与搜索-排序链表/normal"
    normal "myleetcode/排序与搜索-排序链表/normal2"
)

func main() {
    nums := []int{4,3,2,6,5,7,1}
    print(normal.SortList(makeChain(nums)))
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