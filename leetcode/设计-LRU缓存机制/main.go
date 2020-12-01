package main

import (
    "fmt"
    "myleetcode/设计-LRU缓存机制/normal"
)

func main() {
    cache := normal.Constructor(2)
    cache.Put(1,1)
    cache.Put(2,2)
    fmt.Println(cache.Get(1))
    cache.Put(3,3)
    fmt.Println(cache.Get(2))
    cache.Put(4,4)
    fmt.Println(cache.Get(1))
    fmt.Println(cache.Get(3))
    fmt.Println(cache.Get(4))
}
