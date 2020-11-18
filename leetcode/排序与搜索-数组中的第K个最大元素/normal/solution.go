package normal

func FindKthLargest(nums []int, k int) int {
    kmax := make([]int, 0, k)
    for _, num := range nums {
        kmax = push(kmax, num, k)
    }

    return kmax[len(kmax)-1]
}

func push(kmax []int, n, k int) []int {
    count := len(kmax)

    for i := 0; i < count; i ++ {
        if n > kmax[i] {
            tmp := kmax[i]
            kmax[i] = n
            n = tmp
        }
    }

    if count < k {
        kmax = append(kmax, n)
    }

    return kmax
}