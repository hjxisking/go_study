package main

import (
    "fmt"
    "myleetcode/链表-回文链表/normal"
)

func main() {
    //nums := []int{1,2,2,1}
    //nums := []int{1,2,3,2,1}
    nums := []int{1,2,3,4,2,1}
    chain := makeChain(nums)
    fmt.Println(normal.IsPalindrome(chain))
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