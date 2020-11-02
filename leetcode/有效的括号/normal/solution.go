package normal

func IsValid(s string) bool {
    brackets := map[byte]int{
        '(' : -1,
        ')' : 1,
        '{' : -1,
        '}' : 1,
        '[' : -1,
        ']' : 1,
    }

    bracketPairs := map[byte]byte{
        '(' : ')',
        ')' : '(',
        '{' : '}',
        '}' : '{',
        '[' : ']',
        ']' : '[',
    }

    flag := false
    stack := make([]byte, 0)

    for _, ch := range []byte(s) {
        if b, ok := brackets[ch]; ok {
            if b == -1 {   // left bracket, push
                flag = true
                stack = append(stack, ch)
            } else {      // right bracket, pop
                if len(stack) == 0 {
                    return false
                }

                bracket := stack[len(stack)-1]
                stack = stack[:len(stack)-1]

                if bracketPairs[ch] != bracket {
                    return false
                }
            }
        }
    }

    if flag && len(stack) == 0 {
        return true
    } else {
        return false
    }
}
