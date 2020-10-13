package main

import (
    "fmt"
    "math/rand"
    v1 "skiplist/v1"
    "strconv"
    "time"
)

func init() {
    rand.Seed(time.Now().UnixNano())
}

func main() {
    skiplist := v1.NewSkipList(8)
    num := 99
    for i:=0; i<num; i++ {
        key := rand.Intn(num)
        skiplist.Add(key, strconv.Itoa(key))
    }

    key := 8
    skiplist.Add(key, strconv.Itoa(key))

    fmt.Printf("length: %d\n", skiplist.Length())
    fmt.Printf("Search: %v\n", skiplist.Search(key))
    skiplist.Remove(key)
    fmt.Printf("Remove: %d, Search: %v\n", key, skiplist.Search(key))

    skiplist.Print(2)
}
