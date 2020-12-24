package main

import (
    "fmt"
    "myleetcode/链表-奇偶链表/normal"
)

func main() {
    nums := []int{2,1,3,5,6,4,7}
    //nums := []int{1,2,3,4,5,6,7,8}
    chain := makeChain(nums)
    printChain(normal.OddEvenList(chain))
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

func printChain(node *normal.ListNode) {
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