package normal

import (
    "strings"
)

// 回溯算法
func WordBreak(s string, wordDict []string) []string {
    wordDictMap := map[string]bool{}
    for _, word := range wordDict {
        wordDictMap[word] = true
    }

    ans := []string{}
    cache := map[string][][]string{}
    wordList := backtracking(s, wordDictMap, cache)
    for _, words := range wordList {
        ans = append(ans, strings.Join(words, " "))
    }

    return ans
}

func backtracking(s string, wordDictMap map[string]bool, cache map[string][][]string) [][]string {
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
                if words, ok := cache[newS]; !ok {
                    words = backtracking(newS, wordDictMap, cache)
                    cache[newS] = words
                    ans = append(ans, combine(word, words)...)
                } else {
                    ans = append(ans, combine(word, words)...)
                }
            }
        }
    }
    return ans
}

func combine(s string, strs [][]string) [][]string{
    res := [][]string{}
    for i := 0; i < len(strs); i ++ {
        res = append(res, append([]string{s}, strs[i]...))
    }

    return res
}

func check(s string, wordDictMap map[string]bool) bool {
    _, ok := wordDictMap[s]
    return ok
}