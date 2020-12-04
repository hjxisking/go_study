package normal

// 左下角开始搜索，如果目标比当前大，则向上减小，反之则继续向右，直到出边界
func SearchMatrix(matrix [][]int, target int) bool {
    x := len(matrix) - 1
    y := 0

    for x >= 0 && y <= len(matrix[0])-1 {
        if matrix[x][y] == target {
            return true
        } else if matrix[x][y] > target {
            x --
        } else {
            y ++
        }
    }

    return false
}