package heapMax

/**
 * 堆的特性
 * 1. i节点的父节点：(i-1)/2
 * 2. i节点的左右子节点：2*i+1, 2*i+2
 */
type heapMax struct {
    node []int
}

func (h *heapMax) init() {
    for i := len(h.node) - 1; i >= 0; i -- {
        parent := h.parent(i)
        if h.value(i) > h.value(parent) {
            h.swap(parent, i)
            h.sort(i)
        }
    }
}

func (h *heapMax) swap(i, j int) {
    tmp := h.value(j)
    h.set(j, h.value(i))
    h.set(i, tmp)
}

func (h *heapMax) sort(i int) {
    left := h.leftChild(i)
    right := h.rightChild(i)

    exchange := -1

    if right != -1 {    // 有右节点
        if h.value(left) >= h.value(right) && h.value(i) < h.value(left) {
            exchange = left
        }

        if h.value(right) > h.value(left) && h.value(i) < h.value(right) {
            exchange = right
        }
    } else if left != -1 {  // 没有右节点，但是有左节点
        if h.value(left) > h.value(i) {
            exchange = left
        }
    }

    if exchange != -1 {
        h.swap(i, exchange)
        h.sort(exchange)
    }
}

func (h *heapMax) parent(i int) int {
    return (i - 1) / 2
}

func (h *heapMax) leftChild(i int) int {
    left := 2 * i + 1
    if left > len(h.node) - 1 {
        left = -1
    }
    return left
}

func (h *heapMax) rightChild(i int) int {
    right := 2 * i + 2
    if right > len(h.node) - 1 {
        right = -1
    }
    return right
}

func (h *heapMax) value(i int) int {
    return h.node[i]
}

func (h *heapMax) set(i int, value int) {
    h.node[i] = value
}

func (h *heapMax) remove() int {
    max := h.value(0)
    h.set(0, h.value(len(h.node) - 1))
    h.node = h.node[:len(h.node)-1]
    h.sort(0)

    return max
}

func newHeapMax(node []int) *heapMax {
    heap := &heapMax{node:node}
    heap.init()
    return heap
}

func FindKthLargest(nums []int, k int) int {
    ans := -1
    heap := newHeapMax(nums)
    for i := 0; i < k; i ++ {
        ans = heap.remove()
    }

    return ans
}
