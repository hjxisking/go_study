package normal

import "sync"

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */

// hashMap和双向链表

type node struct {
    key int
    value int
    next *node
    prev *node
}

type LRUCache struct {
    cache map[int]*node
    size int
    cap int
    head *node
    tail *node
    mu sync.Mutex
}


func Constructor(capacity int) LRUCache {
    head := &node{}
    tail := &node{}
    head.next = tail
    tail.prev = head

    return LRUCache{
        cache : map[int]*node{},
        size : 0,
        cap  : capacity,
        head : head,
        tail : tail,
    }
}


func (this *LRUCache) Get(key int) int {
    if cache, ok := this.cache[key]; ok {
        this.mu.Lock()
        defer this.mu.Unlock()

        this.setTop(cache)
        return cache.value
    } else {
        return -1
    }
}


func (this *LRUCache) Put(key int, value int)  {
    node := &node{
        key:   key,
        value: value,
        next:  nil,
        prev:  nil,
    }

    this.mu.Lock()
    defer this.mu.Unlock()

    if cache, ok := this.cache[key]; ok {
        cache.value = value
        this.setTop(cache)
    } else {
        if this.size < this.cap {
            this.size ++
        } else {
            this.deleteTail()
        }

        this.setTop(node)
        this.cache[key] = node
    }

}

func (this *LRUCache) setTop(cache *node) {
    if cache.next != nil && cache.prev != nil { // 已经存在的节点，需要从原来的prev和next链接中断开出来
        next := cache.next
        prev := cache.prev
        next.prev = prev
        prev.next = next
    }

    top := this.head.next
    this.head.next = cache
    cache.prev = this.head
    cache.next = top
    top.prev = cache
}

func (this *LRUCache) deleteTail() {
    node := this.tail.prev
    if node != this.head {
        this.tail.prev = node.prev
        node.prev.next = this.tail
        delete(this.cache, node.key)
    }
}