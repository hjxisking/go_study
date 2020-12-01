package normal

func MaxProfit(prices []int) int {
    in := 0
    max := 0

    for i := 1; i < len(prices); i ++ {
        profit := prices[i] - prices[in]
        if profit > 0 {
            if profit > max {
                max = profit
            }
        } else {
            in = i
        }
    }

    return max
}