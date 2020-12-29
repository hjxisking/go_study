package normal

import (
    "math/rand"
    "time"
)

type RandomizedSet struct {
    index []int
    data map[int]int
}


/** Initialize your data structure here. */
func Constructor() RandomizedSet {
    return RandomizedSet{
        index: []int{},
        data:  map[int]int{},
    }
}


/** Inserts a value to the set. Returns true if the set did not already contain the specified element. */
func (this *RandomizedSet) Insert(val int) bool {
    if _, ok := this.data[val]; ok {
        return false
    } else {
        this.index = append(this.index, val)
        this.data[val] = len(this.index) - 1
        return true
    }
}


/** Removes a value from the set. Returns true if the set contained the specified element. */
// 通过this.data找到将要删除的index下标i，然后将该index下标i的内容由index数组最后一个值覆盖，再删除最后那个数组值，同时修正原来末尾的val对应的hashmap里存储的下标为i
func (this *RandomizedSet) Remove(val int) bool {
    if i, ok := this.data[val]; ok {
        lastVal := this.index[len(this.index)-1]
        this.index[i] = lastVal
        this.data[lastVal] = i
        delete(this.data, val)
        this.index = this.index[:len(this.index)-1]
        return true
    }
    return false
}


/** Get a random element from the set. */
func (this *RandomizedSet) GetRandom() int {
    rand.Seed(time.Now().UnixNano())
    random := rand.Intn(len(this.index))
    return this.index[random]
}


/**
 * Your RandomizedSet object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Insert(val);
 * param_2 := obj.Remove(val);
 * param_3 := obj.GetRandom();
 */