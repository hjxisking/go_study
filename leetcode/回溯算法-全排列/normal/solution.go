package normal

func Permute(nums []int) [][]int {
    if len(nums) == 0 {
        return [][]int{}
    }
    if len(nums) == 1 {
        return [][]int{nums}
    }

    res := make([][]int, 0)
    for i := 0; i < len(nums); i ++ {
        num := nums[i]
        tmpNums := cut(nums, i)
        tmpAppend := Permute(tmpNums)
        res = append(res, multiAppend(num, tmpAppend)...)
    }
    return res
}

func cut(nums []int, i int) []int {
    res := make([]int, 0, len(nums)-1)
    for idx, num := range nums {
        if idx != i {
            res = append(res, num)
        }
    }
    return res
}

func multiAppend(num int, sets [][]int) [][]int {
    multiSets := make([][]int, 0)
    for _, set := range sets {
        if len(set) == 0 {
            multiSets = append(multiSets, []int{num})
        } else {
            multiSets = append(multiSets, append(append([]int{}, num), set...))
        }
    }

    return multiSets
}