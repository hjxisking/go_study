package normal

func LongestPalindrome(s string) string {
    palindrome := ""
    strMap := map[int32][]int{}
    for i, c := range s {
        strMap[c] = append(strMap[c], i)
    }

    for c, idxs := range strMap {
        if len(idxs) > 1 {  // 字母只出现一次，是不可能成为回文的起始字母的
            for i:=0; i<len(idxs)-1; i++ {
                start := idxs[i]
                for end := idxs[len(idxs)-1]; end > 0; end -- {
                    if end - start + 1 <= len(palindrome) { // 待判断的回文长度已经小于已存在的回文长度，就不用计算了
                        break
                    }

                    candidate := s[start:end+1]
                    if isPalindrome(candidate) {
                        palindrome = candidate
                    }
                }
            }
        } else if palindrome == "" {
            palindrome = string(c)
        }
    }

    return palindrome
}

func isPalindrome(s string) bool {
    i, j := 0, len(s)-1
    for {
        if i > j {
            return true
        }

        if s[i] != s[j] {
            return false
        }

        i ++
        j --
    }
}