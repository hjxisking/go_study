package main

import (
    "fmt"
    "myleetcode/数学与数字-分数到小数/normal"
)

func main() {
    //numerator := 1
    //denominator := 2

    //numerator := 1
    //denominator := 3

    // 0.(012)
    //numerator := 4
    //denominator := 333

    // 0.(16)
    //numerator := 1
    //denominator := 6

    // -6.25
    //numerator := -50
    //denominator := 8

    // 11
    //numerator := -22
    //denominator := -2

    // -0.58(3)
    numerator := 7
    denominator := -12

    fmt.Println(normal.FractionToDecimal(numerator, denominator))
}
