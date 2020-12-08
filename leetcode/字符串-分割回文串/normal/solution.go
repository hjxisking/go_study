package normal

func Partition(s string) [][]string {
    return backtracking(s)
}

func backtracking(s string) [][]string {
    res := [][]string{}
    if len(s) == 1 {
        return append(res, []string{s})
    }

    for i := 1; i <= len(s); i ++ {
        preString := s[:i]
        postString := s[i:]
        if isPalindrome(preString) {
            if len(postString) == 0 {
                res = append(res, []string{preString})
            } else {
                res = append(res, combine(preString, backtracking(postString))...)
            }
        }
    }

    return res
}

func combine(s string, subs [][]string) [][]string {
    for i, sub := range subs {
        subs[i] = append([]string{s}, sub...)
    }

    return subs
}

func isPalindrome(s string) bool {
    i := 0
    j := len(s) - 1

    for i <= j {
        if s[i] == s[j] {
            i ++
            j --
        } else {
            return false
        }
    }

    return true
}