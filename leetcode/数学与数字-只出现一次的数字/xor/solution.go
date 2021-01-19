package xor

func SingleNumber(nums []int) int {
    ans := 0
    for _, n := range nums {
        ans ^= n
    }

    return ans
}