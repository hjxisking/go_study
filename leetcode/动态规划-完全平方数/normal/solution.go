package normal

import "math"

func NumSquares(n int) int {
    cache := map[int]int{}
    return recursion(n, cache)
}

func recursion(n int, cache map[int]int) int {
    ans := 1 << 63 - 1
    nearestSquare := int(math.Sqrt(float64(n)))
    if nearestSquare * nearestSquare == n {
        return 1
    }

    for ; nearestSquare > 0; nearestSquare -- {
        var num int
        next := n-nearestSquare*nearestSquare
        if ns, ok := cache[next]; ok {
            num = 1 + ns
        } else {
            ns = recursion(next, cache)
            num = 1 + ns
            cache[next] = ns
        }
        if num < ans {
            ans = num
        }
    }
    return ans
}