package normal

type MedianFinder struct {
    min *heapMin
    max *heapMax
    size int
}


/** initialize your data structure here. */
func Constructor() MedianFinder {
    return MedianFinder{
        min : NewHeapMin(),
        max : NewHeapMax(),
    }
}


func (this *MedianFinder) AddNum(num int)  {
    if this.max.Len() == 0 {
        this.max.Push(num)
    } else {
        max := this.max.Top()

        if num < max {
            this.max.Push(num)
        } else {
            this.min.Push(num)
        }

        this.balance()
    }
    this.size ++
}


func (this *MedianFinder) FindMedian() float64 {
    if this.size % 2 == 1 {
        if this.max.Len() > this.min.Len() {
            return float64(this.max.Top())
        } else {
            return float64(this.min.Top())
        }
    } else {
        return float64(this.max.Top() + this.min.Top()) / 2
    }
}

func (this *MedianFinder) balance() {
    maxLen := this.max.Len()
    minLen := this.min.Len()
    if maxLen > minLen && maxLen - minLen > 1 {
        this.min.Push(this.max.Pop())
    } else if minLen > maxLen && minLen - maxLen > 1 {
        this.max.Push(this.min.Pop())
    }
}


/**
 * Your MedianFinder object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AddNum(num);
 * param_2 := obj.FindMedian();
 */