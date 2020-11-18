package main

import (
    "fmt"
    "myleetcode/数学与数字-回文数/normal"
    revert "myleetcode/数学与数字-回文数/反转一半数字"
)

func main() {
    //num := 121
    //num := -121
    //num := 0
    //num := 123
    num := 12321

    fmt.Println(normal.IsPalindrome(num))
    fmt.Println(revert.IsPalindrome(num))
}
