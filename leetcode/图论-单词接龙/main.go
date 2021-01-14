package main

import (
    "fmt"
    "myleetcode/图论-单词接龙/normal"
)

func main() {
    beginWord := "hit"
    endWord := "cog"
    wordList := []string{"hot", "dot", "dog", "lot", "log", "cog"}

    fmt.Println(normal.LadderLength(beginWord, endWord, wordList))
}
