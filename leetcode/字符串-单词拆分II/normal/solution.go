package normal

import "strings"

// 回溯算法
func WordBreak(s string, wordDict []string) []string {
    wordDictMap := map[string]bool{}
    for _, word := range wordDict {
        wordDictMap[word] = true
    }

    ans := []string{}
    wordList := backtracking(s, wordDictMap)
    for _, words := range wordList {
        ans = append(ans, strings.Join(words, " "))
    }

    return ans
}

func backtracking(s string, wordDictMap map[string]bool) [][]string {
    if len(s) == 1 {
        if check(s, wordDictMap) {
            return [][]string{[]string{s}}
        } else {
            return [][]string{}
        }
    }

    ans := [][]string{}
    for i := 1; i <= len(s); i ++ {
        word := s[:i]
        if check(word, wordDictMap) {
            newS := s[i:]
            if newS == "" {
                ans = append(ans, []string{word})
            } else {
                words := backtracking(newS, wordDictMap)
                ans = append(ans, combine(word, words)...)
            }
        }
    }
    return ans
}

func combine(s string, strs [][]string) [][]string{
    for i := 0; i < len(strs); i ++ {
        strs[i] = append([]string{s}, strs[i]...)
    }

    return strs
}

func check(s string, wordDictMap map[string]bool) bool {
    _, ok := wordDictMap[s]
    return ok
}