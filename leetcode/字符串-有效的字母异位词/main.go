package main

import (
    "fmt"
    "myleetcode/字符串-有效的字母异位词/buildInSort"
    "myleetcode/字符串-有效的字母异位词/normal"
)

func main() {
    s := "anagram"
    t := "nagaram"

    //s := "cat"
    //t := "rat"

    fmt.Println(normal.IsAnagram(s, t))
    fmt.Println(buildInSort.IsAnagram(s, t))
}
