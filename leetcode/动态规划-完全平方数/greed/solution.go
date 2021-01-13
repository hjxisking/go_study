package greed

import "math"

func NumSquares(n int) int {
    squares := []int{}
    for i := 1; i * i <= n; i ++ {
        squares = append(squares, i * i)
    }

    for i := 1; i <= n; i ++ {
        if isDividedBy(n, i, squares) {
            return i
        }
    }
    return -1
}

// n 是否可以被分割成 k 个完全平方数
func isDividedBy(n, k int, squares []int) bool {
    if k == 1 {
        square := int(math.Sqrt(float64(n)))
        return n == square * square
    }

    for _, square := range squares {
        if square < n {
            if isDividedBy(n-square, k - 1, squares) {
                return true
            }
        }
    }
    return false
}