package hash

func Intersect(nums1 []int, nums2 []int) []int {
    ans := []int{}
    hashMap := map[int]int{}
    n1 := len(nums1)
    n2 := len(nums2)
    var first []int
    var second []int
    if n1 < n2 {
        first = nums1
        second = nums2
    } else {
        first = nums1
        second = nums2
    }

    for _, num := range first {
        if _, ok := hashMap[num]; !ok {
            hashMap[num] = 0
        }
        hashMap[num] ++
    }

    for _, num := range second {
        if value, ok := hashMap[num]; ok && value > 0 {
            hashMap[num] --
            ans = append(ans, num)
        }
    }

    return ans
}