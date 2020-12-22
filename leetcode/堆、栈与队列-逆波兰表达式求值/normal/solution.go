package normal

import "strconv"

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

func newStack() *stack {
    return &stack{data:[]int{}}
}

func isOperator(s string) bool {
    return s == "+" || s == "-" || s == "*" || s == "/"
}

func calc(i, j int, op string) int {
    res := 0
    switch op {
    case "+" :
        res = i + j
    case "-":
        res = i - j
    case "*":
        res = i * j
    case "/":
        res = i / j
    }
    return res
}

func EvalRPN(tokens []string) int {
    st := newStack()
    for _, token := range tokens {
        if isOperator(token) {
            last := st.pop()
            prev := st.pop()
            st.push(calc(prev, last, token))
        } else {
            num, _ := strconv.Atoi(token)
            st.push(num)
        }
    }
    return st.pop()
}