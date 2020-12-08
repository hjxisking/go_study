package main

import (
    "fmt"
    "myleetcode/字符串-单词拆分/normal"
)

func main() {
    //s := "applepenapple"
    //wordDict := []string{"apple", "pen"}
    s := "catsandog"
    wordDict := []string{"cats", "dog", "sand", "and", "cat"}
    fmt.Println(normal.WordBreak(s, wordDict))
}
