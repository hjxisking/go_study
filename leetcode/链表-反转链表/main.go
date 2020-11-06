package main

import (
    "fmt"
    //rev "myleetcode/链表-反转链表/loop"
    rev "myleetcode/链表-反转链表/recursion"
)

func main() {
    nodes := []int{1,2,3,4,5}
    head := &rev.ListNode{
        Val:  nodes[0],
        Next: nil,
    }
    current := head
    for i := 1; i < len(nodes); i ++ {
        node := &rev.ListNode{
            Val:  nodes[i],
            Next: nil,
        }
        current.Next = node
        current = node
    }

    print(head)

    current = rev.ReverseList(head)

    print(current)
}

func print(node *rev.ListNode) {
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
