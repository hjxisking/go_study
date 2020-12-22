package normal

type stack struct {
    data []interface{}
}

func (s *stack) push(n interface{}) {
    s.data = append(s.data, n)
}

func (s *stack) pop() interface{} {
    res := s.data[len(s.data)-1]
    s.data = s.data[:len(s.data)-1]
    return res
}

func (s *stack) top() interface{} {
    return s.data[len(s.data)-1]
}

func (s *stack) isEmpty() bool {
    return len(s.data) == 0
}

func newStack() *stack {
    return &stack{data:[]interface{}{}}
}

func calcOnce(numStack, opStack *stack) {
    if !numStack.isEmpty() && !opStack.isEmpty() {
        last := numStack.pop().(int)
        prev := numStack.pop().(int)
        op := opStack.pop().(byte)
        numStack.push(operate(prev, last, op))
    }
}

func calc(numStack, opStack *stack) {
    if !numStack.isEmpty() && !opStack.isEmpty() {
        last := numStack.pop().(int)
        for !opStack.isEmpty() {
            op := opStack.pop().(byte)
            if !opStack.isEmpty() {
                prevOp := opStack.top().(byte)
                if prevOp == '-' {
                    if op == '+' {
                        op = '-'
                    } else if op == '-' {
                        op = '+'
                    }
                }
            }
            prev := numStack.pop().(int)
            last = operate(prev, last, op)
        }
        numStack.push(last)
    }
}

// 思路，遇到是操作符，如果上一个操作符是*/，那么将*/计算成结果重新插入栈，最终栈里面只存在+-
// 计算只有+-的栈时，由于是从栈的尾部开始计算，所以在计算当前操作符时需要查看上一个操作符是不是-，如果是的话这次操作取反
func Calculate(s string) int {
    isPrevNum := false
    numStack := newStack()
    opStack := newStack()
    for _, char := range []byte(s) {
        if char == ' ' {
            continue
        } else if isOperate(char) {
            // 当char是操作符时，目的时让操作符栈中*/先处理完，只留下+-
            if !opStack.isEmpty() && isPriorityOperate(opStack.top().(byte)) {    // 消灭栈中*/优先级较高的操作符
                calcOnce(numStack, opStack)
            }

            opStack.push(char)
            isPrevNum = false
        } else {
            num := char2int(char)
            if isPrevNum {
                num = numStack.pop().(int) * 10 + num
            }
            numStack.push(num)
            isPrevNum = true
        }
    }
    calc(numStack, opStack)
    return numStack.top().(int)
}

func isOperate(char byte) bool {
    return char == '+' || char == '-' || char == '*' || char == '/'
}

func isPriorityOperate(char byte) bool {
    return char == '*' || char == '/'
}

func operate(i, j int, op byte) int {
    switch op {
    case '+':
        return i + j
    case '-':
        return i - j
    case '*':
        return i * j
    case '/':
        return i / j
    }
    return 0
}

func char2int(char byte) int {
    return int(char - '0')
}