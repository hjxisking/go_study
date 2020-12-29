package normal

func alpha2int(alpha byte) int {
    return int(alpha - 'A') + 1
}

func TitleToNumber(s string) int {
    ans := 0
    for _, char := range []byte(s) {
        ans = ans * 26 + alpha2int(char)
    }
    return ans
}