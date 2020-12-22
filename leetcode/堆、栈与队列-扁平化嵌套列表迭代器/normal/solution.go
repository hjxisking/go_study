package normal


// This is the interface that allows for creating nested lists.
// You should not implement it, or speculate about its implementation
type NestedInteger struct {
}

// Return true if this NestedInteger holds a single integer, rather than a nested list.
func (n NestedInteger) IsInteger() bool {}

// Return the single integer that this NestedInteger holds, if it holds a single integer
// The result is undefined if this NestedInteger holds a nested list
// So before calling this method, you should have a check
func (n NestedInteger) GetInteger() int {}

// Set this NestedInteger to hold a single integer.
func (n *NestedInteger) SetInteger(value int) {}

// Set this NestedInteger to hold a nested list and adds a nested integer to it.
func (n *NestedInteger) Add(elem NestedInteger) {}

// Return the nested list that this NestedInteger holds, if it holds a nested list
// The list length is zero if this NestedInteger holds a single integer
// You can access NestedInteger's List element directly if you want to modify it
func (n NestedInteger) GetList() []*NestedInteger {}


type NestedIterator struct {
    nums []int
    cursor int
}

func Constructor(nestedList []*NestedInteger) *NestedIterator {
    nums := []int{}

    for _, nestedInteger := range nestedList {
        nums = append(nums, explainNestedInteger(nestedInteger)...)
    }

    nestedIterator := &NestedIterator{
        nums:   nums,
        cursor: -1,
    }
    return nestedIterator
}

func explainNestedInteger(nestedInteger *NestedInteger) []int {
    nums := []int{}
    list := nestedInteger.GetList()
    if len(list) == 0 {
        if nestedInteger.IsInteger(){
            nums = append(nums, nestedInteger.GetInteger())
        }
    } else {
        for _, integer := range list {
            if integer.IsInteger() {
                nums = append(nums, integer.GetInteger())
            } else {
                nums = append(nums, explainNestedInteger(integer)...)
            }
        }
    }

    return nums
}

func (this *NestedIterator) Next() int {
    this.cursor ++
    return this.nums[this.cursor]
}

func (this *NestedIterator) HasNext() bool {
    return this.cursor + 1 < len(this.nums)
}