package normal2

type stack struct {
    data []int
}

func (s *stack) push(n int) {
    s.data = append(s.data, n)
}

func (s *stack) pop() int {
    res := s.data[len(s.data)-1]
    s.data = s.data[:len(s.data)-1]
    return res
}

func (s *stack) top() int {
    return s.data[len(s.data)-1]
}

func (s *stack) isEmpty() bool {
    return len(s.data) == 0
}

func newStack() *stack {
    return &stack{data:[]int{}}
}

func isOperate(char byte) bool {
    return char == '+' || char == '-' || char == '*' || char == '/'
}

func char2int(char byte) int {
    return int(char - '0')
}

func operate(st *stack, op byte) {
    if op == '*' || op == '/' {
        last := st.pop()
        prev := st.pop()
        if op == '*' {
            st.push(prev * last)
        } else {
            st.push(prev / last)
        }
    }
}

// 思路：一个存放有正负数的栈，2-3 = stack.push(+2), stack.push(-3)，最后统计sum
// 过程中遇到*/，优先计算出来重新入栈
func Calculate(s string) int {
    st := newStack()
    isPrevNum := false
    op := byte('#')
    for _, char := range []byte(s) {
        if char == ' ' {
            continue
        } else if isOperate(char) {
            operate(st, op)
            op = char
            isPrevNum = false
        } else {
            num := char2int(char)
            if op == '-' {
                num *= -1
            }
            if isPrevNum {
                num = st.pop() * 10 + num
            }
            st.push(num)
            isPrevNum = true
        }
    }
    operate(st, op)

    ans := 0
    for !st.isEmpty() {
        ans += st.pop()
    }
    return ans
}