package normal

import (
    "strconv"
)

func MaxPoints(points [][]int) int {
    if len(points) == 1 {
        return 1
    }

    ans := 0

    // 2层循环，每次计算以第i个元素为起点，hash统计后续元素和i的斜率计数，并保留最大值，相同坐标单独计算
    for i := 0; i < len(points) - 1; i ++ {
        slopes := map[string]int{}
        max := 1
        duplicate := 0
        for j := i + 1; j < len(points); j ++ {
            x, y := 0, 0
            if points[j][0] >= points[i][0] {
                x = points[j][0] - points[i][0]
                y = points[j][1] - points[i][1]
            } else {
                x = points[i][0] - points[j][0]
                y = points[i][1] - points[j][1]
            }

            if x == 0 && y == 0 {
                duplicate++
                continue
            }

            var slope string
            if x == 0 {
                slope = strconv.Itoa(points[i][0])
            } else if y == 0 {
                slope = strconv.Itoa(points[i][1])
            } else {
                gcd := getGcd(x, y)
                x = x / gcd
                y = y / gcd
                slope = strconv.Itoa(x) + "." + strconv.Itoa(y)
            }
            if _, ok := slopes[slope]; !ok {
                slopes[slope] = 1
            }
            slopes[slope] ++

            if slopes[slope] > max {
                max = slopes[slope]
            }
        }

        if max + duplicate > ans {
            ans = max + duplicate
        }
    }

    return ans
}

// 最大公约数
func getGcd(x, y int) int {
    if x < 0 {
        x = x * -1
    }
    if y < 0 {
        y = y * -1
    }

    if x < y {
        x, y = y, x
    }

    if y == 0 {
        return x
    }

    tmp := x % y
    x = y
    y = tmp
    return getGcd(x, y)
}