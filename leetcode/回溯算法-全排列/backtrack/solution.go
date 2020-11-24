package backtrack

func Permute(nums []int) [][]int {
    res := make([][]int, 0)
    maxDepth := len(nums)
    if maxDepth == 0 {
        return res
    }

    path := make([]int, 0)
    used := make([]bool, len(nums))
    dfs(nums, used, maxDepth, 0, &path, &res)

    return res
}

func dfs(nums []int, used []bool, maxDepth int, depth int, path *[]int, res *[][]int) {
    if maxDepth == depth {
        tmp := make([]int, depth)
        copy(tmp, *path)
        *res = append(*res, tmp)
        return
    }

    for i := 0; i < maxDepth; i ++ {
        if used[i] {
            continue
        } else {
            *path = append(*path, nums[i])
            used[i] = true
            dfs(nums, used, maxDepth, depth + 1, path, res)
            used[i] = false
            *path = (*path)[:len(*path)-1]
        }
    }
}