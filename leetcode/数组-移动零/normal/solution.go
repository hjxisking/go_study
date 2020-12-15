package normal

// 双指针
// 最初i，j同时自增，直到遇到第一个0，i停止自增，j自增，所以只要数组包含0，那么i一定指向值为0的下标，所以当和j交换完之后i自增后仍指向0
func MoveZeroes(nums []int)  {
    n := len(nums)
    i, j := 0, 0

    for j < n {
        if nums[j] != 0 {
            nums[i], nums[j] = nums[j], nums[i]
            i ++
        }
        j ++
    }
}