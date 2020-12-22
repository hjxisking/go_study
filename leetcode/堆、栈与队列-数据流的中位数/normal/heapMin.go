package normal

type heapMin struct {
    data []int
    size int
}

func NewHeapMin() *heapMin {
    return &heapMin{
        data: make([]int, 0),
        size: 0,
    }
}

func (this *heapMin) parent(i int) int {
    return (i - 1) / 2
}

func (this *heapMin) leftChild(i int) int {
    left := 2 * i + 1
    if left > this.size - 1 {
        return -1
    }
    return left
}

func (this *heapMin) rightChild(i int) int {
    right := 2 * i + 2
    if right > this.size - 1 {
        return -1
    }
    return right
}

func (this *heapMin) swap(i, j int) {
    this.data[i], this.data[j] = this.data[j], this.data[i]
}

func (this *heapMin) sort(i int) {
    left := this.leftChild(i)
    right := this.rightChild(i)

    min := this.min(left, right)
    if min != -1 && this.data[i] > this.data[min] {
        this.swap(min, i)
        this.sort(min)
    }
}

func (this *heapMin) min(i, j int) int {
    if i != -1 && j != -1 {
        if this.data[i] < this.data[j] {
            return i
        }
        return j
    } else if i != -1 && j == -1 {
        return i
    } else if i == -1 && j != -1 {
        return j
    } else {
        return -1
    }
}

func (this *heapMin) up(i int) {
    parent := this.parent(i)
    if parent >= 0 && this.data[i] < this.data[parent] {
        this.swap(i, parent)
        this.up(parent)
    }
}

func (this *heapMin) Push(v int) {
    this.data = append(this.data, v)
    this.up(this.size)
    this.size ++
}

func (this *heapMin) Pop() int {
    top := this.data[0]
    this.data[0] = this.data[this.size-1]
    this.data = this.data[:this.size-1]
    this.size --
    this.sort(0)
    return top
}

func (this *heapMin) Len() int {
    return this.size
}

func (this *heapMin) Top() int {
    return this.data[0]
}