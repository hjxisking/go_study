package main

import (
    "fmt"
    "myleetcode/字符串-验证回文/normal"
)

func main() {
    str := "A man, a plan, a canal: Panama"
    //str := "0P"
    fmt.Println(normal.IsPalindrome(str))
}
