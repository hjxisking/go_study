package normal

// 股票走势就像波形图，需要得到最大利润，方法就是在每个波谷买入、波峰抛出，往复一直做下去
func MaxProfit(prices []int) int {
    in := 0
    profit := 0

    for i := 1; i < len(prices); i ++ {
        if prices[i] > prices[in] { // 上升阶段
            if prices[i] < prices[i-1] {    // 过了波峰，i-1是波峰
                in = i
            } else {
                profit += prices[i] - prices[i-1]
                continue
            }
        } else {
            in = i
        }
    }

    return profit
}