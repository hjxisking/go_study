package normal

import "sort"

func WiggleSort(nums []int)  {
    sort.Ints(nums)
    half := len(nums) / 2 + len(nums) % 2
    numsA := make([]int, half)
    numsB := make([]int, len(nums)-half)
    copy(numsA, nums[:half])
    copy(numsB, nums[half:])

    sort.Slice(numsA, func(i, j int) bool {
        return numsA[i] > numsA[j]
    })
    sort.Slice(numsB, func(i, j int) bool {
        return numsB[i] > numsB[j]
    })

    i, j, k := 0, 0, 0
    lenA := len(numsA)
    lenB := len(numsB)
    for k < len(nums) {
        if i < lenA {
            nums[k] = numsA[i]
            i ++
            k ++
        }
        if j < lenB {
            nums[k] = numsB[j]
            j ++
            k ++
        }
    }
}