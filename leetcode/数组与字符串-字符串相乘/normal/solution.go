package normal

import "strconv"

func Multiply(num1 string, num2 string) string {
    n := len(num1)
    m := len(num2)
    ans := "0"

    for i := 0; i < n; i ++ {
        a := num1[i] - '0'
        for j := 0; j < m; j ++ {
            b := num2[j] - '0'
            multiply := int(a*b)
            if multiply != 0 {
                ans = Addition(ans, addPostfixZero(strconv.Itoa(multiply), n - i - 1 + m - j - 1))
            }
        }
    }

    return ans
}

func Addition(num1 string, num2 string) string {
    n := len(num1)
    m := len(num2)

    if n == 0 {
        return num2
    }

    if m == 0 {
        return num1
    }

    max := 0
    if n > m {
        max = n
        num2 = addPrefixZero(num2, n - m)
    } else {
        max = m
        num1 = addPrefixZero(num1, m - n)
    }

    ans := make([]byte, max, max)
    var carry uint8
    for i := max - 1; i >= 0; i -- {
        n1 := num1[i] - '0'
        n2 := num2[i] - '0'

        sum := n1 + n2 + carry
        ans[i] = byte(sum % 10) + '0'
        if sum / 10 > 0 {
            carry = 1
        } else {
            carry = 0
        }
    }

    if carry == 1 {
        tmp := make([]byte, 0, max + 1)
        tmp = append(tmp, '1')
        ans = append(tmp, ans...)
    }

    return string(ans)
}

func addPrefixZero(num string, count int) string {
    res := make([]byte, 0, count + len(num))
    for i := 0; i < count; i ++ {
        res = append(res, '0')
    }

    return string(append(res, []byte(num)...))
}

func addPostfixZero(num string, count int) string {
    res := append(make([]byte, 0, count + len(num)), []byte(num)...)
    for i := 0; i < count; i ++ {
        res = append(res, '0')
    }

    return string(res)
}