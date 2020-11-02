package main

import (
    "fmt"
    "myleetcode/有效的括号/normal"
)

func main() {
    //str := "()[]{}"
    //str := "([])"
    //str := "([a])"
    str := "([{])"
    fmt.Println(normal.IsValid(str))
}
