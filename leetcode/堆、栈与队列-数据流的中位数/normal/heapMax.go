package normal

type heapMax struct {
    data []int
    size int
}

func NewHeapMax() *heapMax {
    return &heapMax{
        data: make([]int, 0),
        size: 0,
    }
}

func (this *heapMax) parent(i int) int {
    return (i - 1) / 2
}

func (this *heapMax) leftChild(i int) int {
    left := 2 * i + 1
    if left > this.size - 1 {
        return -1
    }
    return left
}

func (this *heapMax) rightChild(i int) int {
    right := 2 * i + 2
    if right > this.size - 1 {
        return -1
    }
    return right
}

func (this *heapMax) swap(i, j int) {
    this.data[i], this.data[j] = this.data[j], this.data[i]
}

func (this *heapMax) sort(i int) {
    left := this.leftChild(i)
    right := this.rightChild(i)

    max := this.max(left, right)
    if max != -1 && this.data[i] < this.data[max] {
        this.swap(max, i)
        this.sort(max)
    }
}

func (this *heapMax) max(i, j int) int {
    if i != -1 && j != -1 {
        if this.data[i] > this.data[j] {
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

func (this *heapMax) up(i int) {
    parent := this.parent(i)
    if parent >= 0 && this.data[i] > this.data[parent] {
        this.swap(i, parent)
        this.up(parent)
    }
}

func (this *heapMax) Push(v int) {
    this.data = append(this.data, v)
    this.up(this.size)
    this.size ++
}

func (this *heapMax) Pop() int {
    top := this.data[0]
    this.data[0] = this.data[this.size-1]
    this.data = this.data[:this.size-1]
    this.size --
    this.sort(0)
    return top
}

func (this *heapMax) Len() int {
    return this.size
}

func (this *heapMax) Top() int {
    return this.data[0]
}