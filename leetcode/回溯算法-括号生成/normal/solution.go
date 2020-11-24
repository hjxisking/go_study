package normal

func GenerateParenthesis(n int) []string {
    return generate("", n, 0, 0)
}

func generate(str string, n, left, right int) []string {
    res := []string{}
    if left == n && right == n {
        return append(res, str)
    }

    if left > right {   // 左右括号都可以
        if left < n {
            res = append(res, generate(str+"(", n, left + 1, right)...)
        }
        if right < n {
            res = append(res, generate(str+")", n, left, right + 1)...)
        }
    } else if left == right {   // 只能左括号
        res = append(res, generate(str+"(", n, left + 1, right)...)
    }

    return res
}