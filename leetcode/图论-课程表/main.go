package main

import (
    "fmt"
    "myleetcode/图论-课程表/dfs"
)

func main() {
    //numsCourse := 2
    //prerequisites := [][]int{{1,0}, {0,1}}

    numsCourse := 7
    prerequisites := [][]int{{2,0}, {3,0}, {4, 0}, {0, 1}, {2, 1}, {5, 3}, {4, 5}, {6, 5}}
    fmt.Println(dfs.CanFinish(numsCourse, prerequisites))
}
