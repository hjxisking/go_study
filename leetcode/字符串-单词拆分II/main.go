package main

import (
    "fmt"
    "myleetcode/字符串-单词拆分II/normal"
)

func main() {
    s := "catsanddog"
    wordDict := []string{"cat", "cats", "and", "sand", "dog"}
    res := normal.WordBreak(s, wordDict)
    for _, str := range res {
        fmt.Println(str)
    }
}
