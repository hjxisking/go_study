package normal

import "container/heap"

// dp[i] = min( ...dp[i - coin[0...n]] ) + 1
func CoinChange(coins []int, amount int) int {
    if amount == 0 {
        return 0
    }
    dp := make([]int, amount+1)
    for i := 0; i < len(dp); i ++ {
        dp[i] = -1
    }

    for i := 0; i < len(coins); i ++ {
        if coins[i] <= amount{
            dp[coins[i]] = 1
        }
    }

    for i := 1; i <= amount; i ++ {
        if dp[i] != 1 {
            dpHeap := make(heapMin, 0)
            for j := 0; j < len(coins); j ++ {
                if i - coins[j] > 0 && dp[i - coins[j]] != -1 {
                    heap.Push(&dpHeap, dp[i - coins[j]])
                }
            }
            if dpHeap.Len() > 0 {
                dp[i] = heap.Pop(&dpHeap).(int) + 1
            }
        }
    }

    return dp[amount]
}

type heapMin []int

func (h heapMin) Len() int {
    return len(h)
}

func (h heapMin) Less(i, j int) bool {
    return h[i] < h[j]
}

func (h heapMin) Swap(i, j int) {
    h[i], h[j] = h[j], h[i]
}

func (h *heapMin) Pop() interface{} {
    last := (*h)[len(*h)-1]
    *h = (*h)[0:len(*h)-1]
    return last
}

func (h *heapMin) Push(x interface{}) {
    *h = append(*h, x.(int))
}