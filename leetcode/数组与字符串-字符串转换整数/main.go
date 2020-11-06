package main

import (
    "fmt"
    "myleetcode/数组与字符串-字符串转换整数/normal"
)

func main() {
    //str := "   +91283472332hello1 "
    //str := "9223372036854775808"
    str := "2147483648"
    fmt.Println(normal.MyAtoi(str))
}
