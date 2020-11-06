package normal

import "math"

func MyAtoi(s string) int {
    flag := false
    minus := false
    n := 0
    max := int(math.Pow(2, 31))
    for _, ch := range []byte(s) {
        if ch < '0' || ch > '9' {     // 非数字
            if flag {
                break
            } else {
                if ch == '-' || ch == '+' {        // 数字字符串的开始标记flag还不是true的时候，是允许出现负号的
                    flag = true
                    if ch == '-' {
                        minus = true
                    }
                } else if ch != ' ' {
                    break
                }
            }
        } else {
            flag = true
            ch -= '0'
            if n > max / 10 || ( n == max / 10 && ch > 7) {
                if minus {
                    return max * -1
                } else {
                    return max - 1
                }
            } else {
                n = n * 10 + int(ch)
            }
        }
    }

    if minus {
        n = n * -1
    }

    return n
}
