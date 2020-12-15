package normal

import (
    "math/rand"
    "time"
)

type Solution struct {
    origin []int
    current []int
}


func Constructor(nums []int) Solution {
    return Solution{
        origin : append([]int{}, nums...),
        current : append([]int{}, nums...),
    }
}


/** Resets the array to its original configuration and return it. */
func (this *Solution) Reset() []int {
    for i, num := range this.origin {
        this.current[i] = num
    }
    return this.current
}


/** Returns a random shuffling of the array. */
// Fisher-Yates 洗牌算法
// [9,-4,-1,3] i == 0, random(i, 3) = 2 (范围：0~3)，交换 i=0 和 2 位置 [-1,-4,9,3]，i ++
// [-1,-4,9,3] i == 1, random(i, 3) = 3 (范围：1~3)，交换 i=1 和 3 位置 [-1,3,9,-4]，i ++
// [-1,3,9,-4] i == 2, random(i, 3) = 2 (范围：2~3)，交换 i=2 和 2 位置 [-1,3,9,-4]，i ++
// [-1,3,9,-4] i == 3, random(i, 3) = 3 (范围：3~3)，交换 i=3 和 3 位置 [-1,3,9,-4]，i ++
func (this *Solution) Shuffle() []int {
    n := len(this.current)
    for i := 0; i < n - 1; i ++ {
        key := this.random(i, n)
        this.current[i], this.current[key] = this.current[key], this.current[i]
    }
    return this.current
}

func (this *Solution) random(min, max int) int {
    if min == max {
        return min
    }
    rand.Seed(time.Now().UnixNano())
    return rand.Intn(max-min) + min
}

/**
 * Your Solution object will be instantiated and called as such:
 * obj := Constructor(nums);
 * param_1 := obj.Reset();
 * param_2 := obj.Shuffle();
 */