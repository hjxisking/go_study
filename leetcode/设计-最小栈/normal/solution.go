package normal

import "math"

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */

type node struct {
    value int
    min int
}

type MinStack struct {
    stack []*node
    min int
}


/** initialize your data structure here. */
func Constructor() MinStack {
    return MinStack{
        min : math.MaxInt64,
    }
}


func (this *MinStack) Push(x int)  {
    if x < this.min {
        this.min = x
    }

    data := &node{
        value: x,
        min:   this.min,
    }
    this.stack = append(this.stack, data)
}


func (this *MinStack) Pop()  {
    this.stack = this.stack[:len(this.stack)-1]
    if len(this.stack) > 0 {
        this.min = this.stack[len(this.stack)-1].min
    } else {
        this.min = math.MaxInt64
    }
}


func (this *MinStack) Top() int {
    return this.stack[len(this.stack)-1].value
}


func (this *MinStack) GetMin() int {
    return this.min
}

