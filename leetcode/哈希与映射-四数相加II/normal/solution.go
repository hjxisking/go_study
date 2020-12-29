package normal

func FourSumCount(A []int, B []int, C []int, D []int) int {
    ans := 0
    hashMap := map[int]int{}
    for i := 0; i < len(A); i ++ {
        for j := 0; j < len(B); j ++ {
            key := A[i] + B[j]
            if _, ok := hashMap[key]; !ok {
                hashMap[key] = 0
            }
            hashMap[key] ++
        }
    }

    for k := 0; k < len(C); k ++ {
        for l := 0; l < len(D); l ++ {
            key := (C[k] + D[l]) * -1
            if n, ok := hashMap[key]; ok {
                ans += n
            }
        }
    }

    return ans
}