package main

import (
    "fmt"
    "myleetcode/堆、栈与队列-基本计算器II/normal"
    "myleetcode/堆、栈与队列-基本计算器II/normal2"
)

func main() {
    //calc := "3 + 2 * 2"   // 7
    //calc := "3+5*2-7/2+8" // 18
    //calc := "14/3*2"      // 8
    //calc := "0"           // 0
    calc := "1*2-3/4+5*6-7*8+9/10"  // -24
    fmt.Println(normal.Calculate(calc))
    fmt.Println(normal2.Calculate(calc))
}
