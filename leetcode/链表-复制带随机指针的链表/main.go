package main

import (
    "fmt"
    "myleetcode/链表-复制带随机指针的链表/normal"
)

const NIL = -1

func main() {
    nums := [][]int{{7,NIL}, {13,0}, {11,4}, {10,2}, {1,0}}
    chain := makeChain(nums)
    fmt.Println(chain)
    fmt.Println(normal.CopyRandomList(chain))
}

func makeChain(nums [][]int) *normal.Node {
    hashMap := make([]*normal.Node, len(nums))
    head := &normal.Node{}
    current := head
    for i, num := range nums {
        node := &normal.Node{
            Val:    num[0],
            Next:   nil,
            Random: nil,
        }
        hashMap[i] = node
        current.Next = node
        current = node
    }
    current = head.Next

    for _, num := range nums {
        if num[1] != NIL {
            current.Random = hashMap[num[1]]
        }
        current = current.Next
    }
    return head.Next
}
