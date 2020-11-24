package normal

import "math"

func GrayCode(n int) []int {
    if n == 0 {
        return []int{0}
    }
    if n == 1 {
        return []int{0,1}
    }

    recursion := GrayCode(n-1)
    code := recursion
    for i := len(recursion) - 1; i >= 0; i -- {
        code = append(code, recursion[i] + int(math.Pow(2, float64(n-1))))
    }
    return code
}