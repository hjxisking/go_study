package normal

func FirstUniqChar(s string) int {
    chars := [26]int{}      // 固定大小的数组比动态增加的map效果要好很多
    for _, char := range []byte(s) {
        chars[char-'a'] += 1
    }

    for i, char := range []byte(s) {
        if chars[char-'a'] == 1 {
            return i
        }
    }

    return -1
}