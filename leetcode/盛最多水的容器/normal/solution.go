package normal

import "math"

func MaxArea(height []int) int {
    max := 0
    n := len(height)

    if n == 2 {
        max = int(math.Min(float64(height[0]), float64(height[1])))
    }

    i := 0
    j := n - 1
    for j > i {
        s := int(math.Min(float64(height[i]), float64(height[j]))) * (j - i)

        if height[i] < height[j] {
            i ++
        } else {
            j --
        }

        if s > max {
            max = s
        }
    }

    return max
}