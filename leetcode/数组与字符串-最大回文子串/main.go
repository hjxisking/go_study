package main

import (
    "fmt"
    "myleetcode/数组与字符串-最大回文子串/dynamic"
    "myleetcode/数组与字符串-最大回文子串/normal"
)

func main() {
    //str := "babacabacef"
    //str := "aacabdkacaa"
    //str := "a"
    //str := "cbbd"
    //str := "babadada"
    str := "ac"
    fmt.Println(normal.LongestPalindrome(str))
    fmt.Println(dynamic.LongestPalindrome(str))
}

