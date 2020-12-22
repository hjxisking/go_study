package normal

type heapMin struct {
    data [][]int
    size int
    capacity int
}

func newHeapMin(capacity int) *heapMin {
    return &heapMin{
        data:     make([][]int, 0),
        size:     0,
        capacity: capacity,
    }
}

func (this * heapMin) parent(i int) int {
    return (i - 1) / 2
}

func (this *heapMin) leftChild(i int) int {
    left := 2 * i + 1
    if left < len(this.data) {
        return left
    } else {
        return -1
    }
}

func (this *heapMin) rightChild(i int) int {
    right := 2 * i + 2
    if right < len(this.data) {
        return right
    } else {
        return -1
    }
}

func (this *heapMin) sort(i int) {
    target := this.exchange(i)
    if target != -1 && this.compare(target, i) == -1{
        this.swap(i, target)
        this.sort(target)
    }
}

func (this *heapMin) exchange(i int) int {
    left := this.leftChild(i)
    right := this.rightChild(i)

    if left != -1 && right != -1 {
        if this.compare(left, right) == -1{
            return left
        }
        return right
    } else if left != -1 && right == -1 {
        return left
    } else if left == -1 && right != -1 {
        return right
    } else {
        return -1
    }
}

func (this *heapMin) swap(i, j int) {
    this.data[i], this.data[j] = this.data[j], this.data[i]
}

func (this *heapMin) up(i int) {
    parent := this.parent(i)
    if parent >= 0 && this.compare(i, parent) == -1 {
        this.swap(i, parent)
        this.up(parent)
    }
}

func (this *heapMin) add(k, v int) {
    if this.size < this.capacity {
        this.append(k, v)
        this.up(this.size)
        this.size ++
    } else {
        _, value := this.get(0)
        if v > value {
            this.set(0, k, v)
            this.sort(0)
        }
    }
}

func (this *heapMin) compare(i, j int) int {
    if this.data[i][1] < this.data[j][1] {
        return -1
    } else if this.data[i][1] == this.data[j][1] {
        return 0
    } else {
        return 1
    }
}

func (this *heapMin) set(i, k, v int) {
    this.data[i] = []int{k, v}
}

func (this *heapMin) append(k, v int) {
    this.data = append(this.data, []int{k, v})
}

func (this *heapMin) get(i int) (int, int) {
    return this.data[i][0], this.data[i][1]
}

func TopKFrequent(nums []int, k int) []int {
    ans := []int{}
    hashMap := map[int]int{}
    for _, num := range nums {
        if _, ok := hashMap[num]; !ok {
            hashMap[num] = 0
        }
        hashMap[num] ++
    }

    heap := newHeapMin(k)
    for key, value := range hashMap {
        heap.add(key, value)
    }

    for i := k - 1; i >= 0; i -- {
        key, _ := heap.get(i)
        ans = append(ans, key)
    }

    return ans
}