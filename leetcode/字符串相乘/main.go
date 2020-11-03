package main

import (
    "fmt"
    "myleetcode/字符串相乘/normal"
    "myleetcode/字符串相乘/乘法思路"
)

func main() {
    //num1 := "7234"
    //num2 := "4567"
    //num2 := "0"

    num1 := "7629325891502811763742926406289264396895057"
    num2 := "3872531151515609491056217197663171"

    //fmt.Println(normal.Addition(num1, num2))
    fmt.Println(normal.Multiply(num1, num2))
    fmt.Println(乘法思路.Multiply(num1, num2))
}