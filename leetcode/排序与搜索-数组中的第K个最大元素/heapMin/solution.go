package heapMin

type heapMin struct {
    node []int
    size int
    capacity int
}

func (h *heapMin) up(i int) {
    parent := h.parent(i)
    if parent >= 0 && h.value(i) < h.value(parent) {
        h.swap(i, parent)
        h.sort(i)
        h.up(parent)
    }
}

func (h *heapMin) swap(i, j int) {
    tmp := h.value(j)
    h.set(j, h.value(i))
    h.set(i, tmp)
}

func (h *heapMin) sort(i int) {
    left := h.leftChild(i)
    right := h.rightChild(i)

    exchange := -1

    if right != -1 {    // 有右节点
        if h.value(left) <= h.value(right) && h.value(i) > h.value(left) {
            exchange = left
        }

        if h.value(right) < h.value(left) && h.value(i) > h.value(right) {
            exchange = right
        }
    } else if left != -1 {  // 没有右节点，但是有左节点
        if h.value(left) < h.value(i) {
            exchange = left
        }
    }

    if exchange != -1 {
        h.swap(i, exchange)
        h.sort(exchange)
    }
}

func (h *heapMin) parent(i int) int {
    return (i - 1) / 2
}

func (h *heapMin) leftChild(i int) int {
    left := 2 * i + 1
    if left > len(h.node) - 1 {
        left = -1
    }
    return left
}

func (h *heapMin) rightChild(i int) int {
    right := 2 * i + 2
    if right > len(h.node) - 1 {
        right = -1
    }
    return right
}

func (h *heapMin) value(i int) int {
    return h.node[i]
}

func (h *heapMin) set(i int, n int) {
    h.node[i] = n
}

func (h *heapMin) add(n int) {
    if h.size < h.capacity {
        h.node = append(h.node, n)
        h.size ++
        h.up(h.size-1)
    } else {
        if n > h.value(0) {
            h.set(0, n)
            h.sort(0)
        }
    }
}

func newHeapMin(k int) *heapMin{
    return &heapMin{
        node:     make([]int, 0, k),
        size:     0,
        capacity: k,
    }
}

/**
 * 小根堆的思路就是构建一个k大小的堆，堆顶是最小值
 * 将所有数值都通过k大小的小根堆进行过滤，如果值比堆顶大，则替换之，并且重新恢复小根堆
 * 这样最后就得到了最大的K个值都在这个堆结构中，且最小的那个即第K大就是堆顶
 */
func FindKthLargest(nums []int, k int) int {
    heap := newHeapMin(k)
    for _, n := range nums {
        heap.add(n)
    }

    return heap.value(0)
}