package dynamic

/**
    动态规划方式，https://leetcode-cn.com/problems/longest-palindromic-substring/solution/zui-chang-hui-wen-zi-chuan-by-leetcode-solution/
    已知条件 dp[i][j]：
        1. 当i==j时，即子串长度为1，那么必定为回文，dp[i][j]=true
        2. 当j-i==1时，如果s[i]==s[j]，即子串长度为2，且为相同字符，那么也是回文，dp[i][j]=true
        3. 当j-i==2时，如果s[i]==s[j]，即子串长度为3，且首尾字符相同，那么也是回文，dp[i][j]=true
 */
func LongestPalindrome(s string) string {
    n := len(s)

    if n < 2 {
        return s
    }

    dp := make([][]bool, n)
    palindrome := string(s[0])

    for i := 0; i < n; i ++ {
        dp[i] = make([]bool, n)
    }

    // 重点：先循环j，再循环i，因为如果从i开始循环，那么 0-4 意味着计算整个字符串，但是这时候 1-4 还没计算，所以是无法得知 0-4 的
    // 所以按j开始循环，这样相当于先计算0-1、1-1，再0-2、1-2、2-2，每一步都可以根据之前计算的值来推导条件
    // 动态规划的思想就是，1. 寻找问题的拆解方法，即一个问题的解 = 一个或多个子问题的解，即 f(x) = f(x-1) + ...
    //                    2. 最小子问题一定有一个“常数级”的定义，即 f(1) = 已知解
    // 所以当动态规划不使用递归的方式，而是使用循环的方式时，循环的方向一定是从最小值开始，即 x 从1开始，得到最后的max
    for j := 1; j < n; j ++ {
        for i := 0; i < j; i ++ {
            if i == j {
                dp[i][j] = true
            } else if s[i] == s[j] {
                if j - i < 3 {
                    dp[i][j] = true
                } else {
                    dp[i][j] = dp[i+1][j-1]
                }
            } else {
                dp[i][j] = false
            }

            if dp[i][j] && j - i + 1 > len(palindrome) {
                palindrome = s[i:j+1]
            }
        }
    }

    return palindrome
}
