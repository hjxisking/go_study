package normal

func RemoveDuplicates(nums []int) int {
    n := len(nums)
    i := 0
    for j := 1; j < n; j ++ {
        if nums[i] != nums[j] {
            if j - i > 1 {
                nums[i+1] = nums[j]
            }
            i ++
        }
    }

    return i + 1
}