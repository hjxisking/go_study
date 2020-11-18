package normal

func IsPowerOfTwo(n int) bool {
    if n == 0 {
        return false
    } else {
        return n & (n - 1) == 0
    }
}