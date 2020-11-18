package 反转一半数字

func IsPalindrome(x int) bool {
    if x < 0 || (x % 10 == 0 && x % 10 > 0) {
        return false
    }

    if x == 0 {
        return true
    }

    reverted := 0
    for {
        reverted = reverted * 10 + x % 10
        x = x / 10

        if reverted >= x {
            break
        }
    }

    if reverted == x || reverted / 10 == x {
        return true
    } else {
        return false
    }
}
