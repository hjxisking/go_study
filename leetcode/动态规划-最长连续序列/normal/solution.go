package normal

// 1. 将所有数字hashMap去重
// 2. 循环hashMap中的每个数字，寻找连续序列的第一位，条件是不存在 x - 1，但是存在 x + 1
// 3. 找到开始位之后循环开始判断 + 1 后是否存在，直到不存在，并记录最长长度
func LongestConsecutive(nums []int) int {
    ans := 0
    hashMap := map[int]bool{}
    for _, num := range nums {
        hashMap[num] = true
    }

    for num, _ := range hashMap {
        if _, ok := hashMap[num-1]; !ok {
            i := num
            count := 1
            for {
                if _, ok := hashMap[i+1]; !ok {
                    break
                }
                i ++
                count ++
            }
            if count > ans {
                ans = count
            }
        }
    }

    return ans
}