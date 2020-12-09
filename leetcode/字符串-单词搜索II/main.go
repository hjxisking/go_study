package main

import (
    "fmt"
    "myleetcode/字符串-单词搜索II/normal"
)

func main() {
    //board := [][]byte{{'o','a','a','n'}, {'e','t','a','e'}, {'i','h','k','r'}, {'i','f','l','v'}}
    //words := []string{"oath", "oathk", "pea", "eat", "rain"}

    //board := [][]byte{{'a','b'}, {'c','d'}}
    //words := []string{"abcd"}

    //board := [][]byte{{'a','a'}}
    //words := []string{"a"}

    board := [][]byte{{'a','b'}, {'a','a'}}
    words := []string{"aba", "baa", "bab", "aaab", "aaa", "aaaa", "aaba"}

    strs := normal.FindWords(board, words)
    for _, str := range strs {
        fmt.Println(str)
    }
}
