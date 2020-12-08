package normal

// dp[i] = dp[j] && check(a[j..i-1], wordDict)
func WordBreak(s string, wordDict []string) bool {
    wordDictMap := map[string]bool{}
    for _, word := range wordDict {
        wordDictMap[word] = true
    }

    dp := make([]bool, len(s)+1)
    dp[0] = true
    for i := 1; i <= len(s); i ++ {
        for j := 0; j < i; j ++ {
            if dp[j] && check(s[j:i], wordDictMap) {
                dp[i] = true
                break
            }
        }
    }

    return dp[len(s)]
}

func check(str string, wordDictMap map[string]bool) bool {
    _, ok := wordDictMap[str]
    return ok
}