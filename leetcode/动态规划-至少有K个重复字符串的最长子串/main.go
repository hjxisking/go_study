package main

import (
    "fmt"
    "myleetcode/动态规划-至少有K个重复字符串的最长子串/normal"
)

func main() {
    str := "aaabb"
    k := 3

    //str := "ababacb"
    //k := 3

    //str := "aaabbb"
    //k := 3

    //str := "bbaaacbd"
    //k := 3

    fmt.Println(normal.LongestSubstring(str, k))
}
