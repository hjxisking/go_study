package binarySearch

// 二分搜索，从对角线开始搜索下面和右边，搜索下和右的时候分别可以使用二分法来加快搜索，而不是逐个查询
func SearchMatrix(matrix [][]int, target int) bool {
    x := len(matrix)
    y := len(matrix[0])

    for i := 0; i < min(x, y); i ++ {
        if matrix[i][i] == target {
            return true
        }

        if binarySearch(matrix, target, i, true) {
            return true
        }

        if binarySearch(matrix, target, i, false) {
            return true
        }
    }

    return false
}

func binarySearch(matrix [][]int, target int, start int, vertical bool) bool {
    i := start
    var j int
    if vertical {
        j = len(matrix) - 1
    } else {
        j = len(matrix[0]) - 1
    }

    for i < j {
        if vertical {
            if matrix[j][start] == target {
                return true
            } else if matrix[j][start] < target {
                return false
            } else {
                mid := (j + i) / 2
                if matrix[mid][start] == target || matrix[i][start] == target {
                    return true
                } else if matrix[mid][start] < target {
                    i = mid + 1
                } else {
                    j = mid - 1
                }
            }
        } else {
            if matrix[start][j] == target {
                return true
            } else if matrix[start][j] < target {
                return false
            } else {
                mid := (j + i) / 2
                if matrix[start][mid] == target || matrix[start][i] == target {
                    return true
                } else if matrix[start][mid] < target {
                    i = mid + 1
                } else {
                    j = mid - 1
                }
            }
        }
    }

    return false
}

func min(x, y int) int {
    if x < y {
        return x
    } else {
        return y
    }
}