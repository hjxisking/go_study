package main

import (
    "fmt"
    "myleetcode/反转字符串/normal"
)

func main() {
    str := []byte{'h', 'e', 'l', 'l', 'o'}
    fmt.Println(string(str))
    normal.ReverseString(str)
    fmt.Println(string(str))
}
