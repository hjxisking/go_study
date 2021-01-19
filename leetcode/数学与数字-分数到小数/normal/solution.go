package normal

import "strconv"

func FractionToDecimal(numerator int, denominator int) string {
    if numerator == 0 {
        return "0"
    }

    minus := false
    if (numerator < 0 && denominator > 0) || (numerator > 0 && denominator < 0) {
        minus = true
    }

    if numerator < 0 {
        numerator *= -1
    }

    if denominator < 0 {
        denominator *= -1
    }

    ans := []byte{}
    res := ""
    point := false
    index := 1  // 小数点后的位置，用来对应本次操作时分子和小数位的对应关系
    periodic := false
    decimal := map[int]int{}

    for numerator != 0 {
        appendRes := ""
        newNumerator := numerator

        if numerator < denominator {
            if res == "" {
                res = "0"
            }

            numerator *= 10

            if !point {
                res += "."
                point = true
            }
        }
        division := numerator/denominator
        appendRes = strconv.Itoa(division)
        newNumerator = numerator - division * denominator


        if point {
            if _, ok := decimal[numerator]; !ok {
                decimal[numerator] = index
            } else {
                periodic = true
                break
            }
            index ++
        }
        numerator = newNumerator
        res += appendRes
    }

    pointPos := 1
    for _, char := range []byte(res) {
        if char == '.' {
            break
        } else {
            pointPos ++
        }
    }

    periodicIndex := decimal[numerator]
    if periodic {
        tmp := []byte(res)
        ans = append(ans, tmp[0:pointPos+periodicIndex-1]...)
        ans = append(ans, '(')
        ans = append(ans, tmp[pointPos+periodicIndex-1:]...)
        ans = append(ans, ')')
    } else {
        ans = []byte(res)
    }

    if minus {
        ans = append([]byte{'-'}, ans...)
    }

    return string(ans)
}