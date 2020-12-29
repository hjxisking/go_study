package main

import (
    "fmt"
    "myleetcode/哈希与映射-常数时间插入、删除和获取随机元素/normal"
)

func main() {
    randomSet := normal.Constructor()
    fmt.Println(randomSet.Insert(1))
    fmt.Println(randomSet.Remove(2))
    fmt.Println(randomSet.Insert(2))
    fmt.Println(randomSet.GetRandom())
    fmt.Println(randomSet.Remove(1))
    fmt.Println(randomSet.Insert(2))
    fmt.Println(randomSet.GetRandom())
}
