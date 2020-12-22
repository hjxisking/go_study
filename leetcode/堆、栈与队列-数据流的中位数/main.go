package main

import (
    "fmt"
    "myleetcode/堆、栈与队列-数据流的中位数/normal"
)

func main() {
    //heapMin := normal.NewHeapMin()
    //heapMin.Push(9)
    //heapMin.Push(5)
    //heapMin.Push(1)
    //heapMin.Push(7)
    //heapMin.Push(4)
    //heapMin.Push(2)
    //heapMin.Push(3)
    //heapMin.Push(6)
    //for i := 0; i < 8; i ++ {
    //    fmt.Println(heapMin.Pop())
    //}

    medianFinder := normal.Constructor()
    medianFinder.AddNum(1)
    medianFinder.AddNum(2)
    fmt.Println(medianFinder.FindMedian())
    medianFinder.AddNum(3)
    fmt.Println(medianFinder.FindMedian())
}
