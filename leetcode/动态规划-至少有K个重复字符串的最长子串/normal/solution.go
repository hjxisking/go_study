package normal

// 1. 收集所有出现的字符的次数，不一定要用map，可以用固定长度为26的数组 index = char - 'a'
// 2. 循环s，直到遇到的字符 X 出现次数是少于K的
// 3. 以 X 为界，递归左右2个子串，并比较返回max
// 4. 递归结束的条件是，子串中所有字符的出现次数均大于K，如果字符串长度本身小于K了，就不用分割了
func LongestSubstring(s string, k int) int {
    if len(s) < k {
        return 0
    }

    sCount := [26]int{}
    for _, char := range []byte(s) {
        index := char - 'a'
        sCount[index] ++
    }

    for i, char := range []byte(s) {
        index := char - 'a'
        if sCount[index] < k {
            return max(LongestSubstring(s[:i], k), LongestSubstring(s[i+1:], k))
        }
    }

    // s在递归过程中不断缩小，如果没有在for循环中提前return，就说明本s字符串内的所有字符出现次数都是大于K的，返回其个数即可
    return len(s)
}

func max(x, y int) int {
    if x > y {
        return x
    }
    return y
}
