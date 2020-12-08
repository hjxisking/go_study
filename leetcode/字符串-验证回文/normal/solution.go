package normal

func IsPalindrome(s string) bool {
    i := 0
    j := len(s) - 1

    for i <= j {
        if !isChar(s[i]) {
            i ++
            continue
        }

        if !isChar(s[j]) {
            j --
            continue
        }

        if isEqual(s[i], s[j]) {
            i ++
            j --
        } else {
            return false
        }
    }

    return true
}

func isNum(c byte) bool {
    return c >= 48 && c <= 57
}

func isUpper(c byte) bool {
    return c >= 65 && c <= 90
}

func isLower(c byte) bool {
    return c >= 97 && c <= 122
}

func isChar(c byte) bool {
    if isNum(c) || isUpper(c) || isLower(c) {
        return true
    } else {
        return false
    }
}

func isEqual(c1, c2 byte) bool {
    if isNum(c1) && isNum(c2) && c1 == c2 {
        return true
    }

    if isUpper(c1) {
        c1 += 32
    }

    if isUpper(c2) {
        c2 += 32
    }

    return c1 == c2
}