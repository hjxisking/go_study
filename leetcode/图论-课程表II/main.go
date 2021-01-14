package main

import (
    "fmt"
    "myleetcode/图论-课程表II/dfs"
)

func main() {
    numsCourse := 4
    prerequisites := [][]int{{1,0}, {2,0}, {3, 1}, {3, 2}}

    //numsCourse := 7
    //prerequisites := [][]int{{2,0}, {3,0}, {4, 0}, {0, 1}, {2, 1}, {5, 3}, {4, 5}, {6, 5}}

    fmt.Println(dfs.FindOrder(numsCourse, prerequisites))
}
