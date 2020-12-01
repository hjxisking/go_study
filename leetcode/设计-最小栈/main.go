package main

import (
    "fmt"
    "myleetcode/设计-最小栈/normal"
)

func main() {
    minStack := normal.Constructor()
    minStack.Push(-2)
    minStack.Push(0)
    minStack.Push(-3)
    fmt.Println(minStack.GetMin())
    minStack.Pop()
    fmt.Println(minStack.Top())
    fmt.Println(minStack.GetMin())
}
