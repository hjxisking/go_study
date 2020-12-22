package deque

type myDeque struct {
    data []int
}

func (d *myDeque) addFirst(n int) {
    d.data = append([]int{n}, d.data...)
}

func (d *myDeque) removeFirst() {
    d.data = d.data[1:]
}

func (d *myDeque) getFirst() int {
    return d.data[0]
}

func (d *myDeque) addLast(n int) {
    d.data = append(d.data, n)
}

func (d *myDeque) removeLast() {
    d.data = d.data[:len(d.data)-1]
}

func (d *myDeque) getLast() int {
    return d.data[len(d.data)-1]
}

func (d *myDeque) len() int {
    return len(d.data)
}

func newDeque() *myDeque {
    return &myDeque{data:[]int{}}
}

func cleanQueue(d *myDeque, nums[]int, i, k int) {
    if d.len() > 0 && d.getFirst() == i - k {
        d.removeFirst()
    }

    for d.len() > 0 && nums[d.getLast()] < nums[i] {
        d.removeLast()
    }
}

func MaxSlidingWindow(nums []int, k int) []int {
    ans := []int{}
    d := newDeque()
    for i := 0; i < len(nums); i ++ {
        cleanQueue(d, nums, i, k)
        d.addLast(i)
        if i - k + 1 >= 0 {
            ans = append(ans, nums[d.getFirst()])
        }
    }
    return ans
}