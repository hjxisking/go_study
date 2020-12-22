package main

import (
    "fmt"
    "myleetcode/堆、栈与队列-逆波兰表达式求值/normal"
)

func main() {
    tokens := []string{"10", "6", "9", "3", "+", "-11", "*", "/", "*", "17", "+", "5", "+"}
    fmt.Println(normal.EvalRPN(tokens))
}
