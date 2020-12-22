package normal

func KthSmallest(matrix [][]int, k int) int {
    row := len(matrix) - 1
    col := len(matrix[0]) - 1
    left := matrix[0][0]
    right := matrix[row][col]

    for left < right {
        mid := (left + right) / 2
        i, j := col, 0
        count := 0
        for i >= 0 && j <= row {
            if matrix[i][j] <= mid {
                j ++
                count += i + 1
            } else {
                i --
            }
        }
        if count >= k {
            right = mid
        } else {
            left = mid + 1
        }
    }
    return left
}
