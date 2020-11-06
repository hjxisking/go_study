package 乘法思路

import (
    "strconv"
    "strings"
)

func Multiply(num1 string, num2 string) string {
    n := len(num1)
    m := len(num2)

    size := n + m

    ans := make([]byte, size, size)
    for idx := 0; idx < size; idx ++ {
        ans[idx] = '0'
    }

    for i := n - 1; i >= 0; i -- {
        a := num1[i] - '0'
        for j := m - 1; j >= 0; j -- {
            var up byte
            b := num2[j] - '0'
            multiply := strconv.Itoa(int(a*b))
            count := len(multiply)
            //fmt.Println(a, b)
            //fmt.Printf("%v%v%v%v\n", strings.Repeat(" ", size-(n-1-i)-(m-1-j)-count), multiply, strings.Repeat("0",n-1-i),strings.Repeat("0",m-1-j))
            for k := 0; k < count; k ++ {
                tmp := ans[i+j+1-k] - '0' + multiply[count-k-1] - '0' + up
                up = tmp / 10
                ans[i+j+1-k] = tmp % 10 + '0'
            }

            if up > 0 {
                ans[i+j+1-count] += up      // 这里不用担心进位问题，数组的这个下标位置是可以存放大于9的byte的
            }
            //fmt.Println(string(ans))
        }
    }

    if string(ans) == strings.Repeat("0", size) {
        return "0"
    }

    if ans[0] == '0' {
        return string(ans[1:])
    } else {
        return string(ans)
    }
}