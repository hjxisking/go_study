package normal

func Subsets(nums []int) [][]int {
    if len(nums) == 0 {
        return [][]int{}
    }
    if len(nums) == 1 {
        return [][]int{
            []int{},
            []int{nums[0]},
        }
    }

    num := nums[0]
    nums = nums[1:]
    tmp := Subsets(nums)
    res := multiAppend(num, tmp)
    return res
}

func multiAppend(num int, sets [][]int) [][]int {
    multiSets := sets
    for _, set := range sets {
        if len(set) == 0 {
            multiSets = append(multiSets, []int{num})
        } else {
            multiSets = append(multiSets, append(append([]int{}, set...), num))
        }
    }

    return multiSets
}