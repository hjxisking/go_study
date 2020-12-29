package main

import (
    "fmt"
    "myleetcode/树-二叉树的序列化与反序列化/normal"
)

func main() {
    //str := "1,2,3,nil,nil,4,5"
    //str := "1,2,3,nil,7,4,5"
    //str := "1,2,3,7,nil,4,5"
    str := "1,2,3,nil,nil,4,5,7"

    ser := normal.Constructor()
    deser := normal.Constructor()

    tree := deser.Deserialize(str)
    newStr := ser.Serialize(tree)
    fmt.Println(str)
    fmt.Println(newStr)
}