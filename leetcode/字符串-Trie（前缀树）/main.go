package main

import (
    "fmt"
    "myleetcode/字符串-Trie（前缀树）/normal"
)

func main() {
    trie := normal.Constructor()

    trie.Insert("apple")
    fmt.Println(trie.Search("apple"))
    fmt.Println(trie.Search("app"))
    fmt.Println(trie.StartsWith("app"))
    trie.Insert("app")
    fmt.Println(trie.Search("app"))

    if trie2, ok := trie.ContainKey('a'); ok {
        fmt.Println(trie2.Search("pp"))
        fmt.Println(trie2.Search("app"))
    }
}
