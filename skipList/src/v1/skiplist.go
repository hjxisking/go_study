package v1

import (
    "fmt"
    "math/rand"
    "strconv"
    "strings"
    "sync"
    "syscall"
    "time"
)

func init() {
    rand.Seed(time.Now().UnixNano())
}

type skipNode struct {
    key int
    data interface{}
    next []*skipNode
}

type SkipList struct {
    length int
    level int
    head *skipNode
    tail *skipNode
    mutex *sync.RWMutex
}

func NewSkipList(level int) *SkipList {
    if level <= 0 || level > 32 {
        level = 32
    }

    skiplist := &SkipList{
        length: 0,
        level:  level,
        head:   &skipNode{
            key: -syscall.INFINITE,
            data: nil,
            next: make([]*skipNode, level, level),
        },
        tail:   &skipNode{
            key:  syscall.INFINITE,
            data: nil,
            next: nil,
        },
        mutex:  &sync.RWMutex{},
    }

    for i:=0; i<level; i++ {
        skiplist.head.next[i] = skiplist.tail
    }

    return skiplist
}

func (list *SkipList) randLevel() int {
    for i:=1; i<=list.level; i++ {
        rnd := rand.Int()
        if rnd % 2 == 0 {
            return i
        }
    }

    return list.level
}

func (list *SkipList) Add(key int, data interface{}) {
    level := list.randLevel()
    curr := list.head

    node := &skipNode{
        key:  key,
        data: data,
        next: make([]*skipNode, level, level),
    }

    update := make([]*skipNode, level, level)
    for i:=level-1; i>=0; i-- {
        for {
            next := curr.next[i]
            if next != list.tail && key > next.key {
                curr = next
            } else if key == next.key {
                return
            } else {
                // 这里只记录第I层的当前节点需要更改next来插入node，目的是又可能在更下面的层会发现要插入的node其实已经存在了，不需要插入，如果这边不是记录而是直接操作插入，则会重复，流程是自上而下的，所以不能提前知道要插入的节点在下面的层是否已经存在
                update[i] = curr
                break
            }
        }
    }

    for i:=level-1; i>=0; i-- {
        curr := update[i]
        next := curr.next[i]
        curr.next[i] = node
        node.next[i] = next
    }

    list.length ++
}

func (list *SkipList) Remove(key int) {
    exists := false
    curr := list.head
    for i:=list.level-1; i>=0; i-- {
        for {
            next := curr.next[i]
            if next == list.tail || key < next.key {
                break
            } else if key == next.key {
                curr.next[i] = next.next[i]
                exists = true
                break
            } else {
                curr = next
            }
        }
    }

    if exists { list.length -- }
}

func (list *SkipList) Search(key int) interface{} {
    curr := list.head
    for i:=list.level-1; i>=0; i-- {
        for {
            next := curr.next[i]
            if next == list.tail || key < next.key {
                break
            } else if key == next.key {
                return next.data
            } else {
                curr = next
            }
        }
    }

    return nil
}

func (list *SkipList) Length() int {
    return list.length
}

func (list *SkipList) Print(width int) {
    nodeIndex := make(map[*skipNode]int, list.length+2) // 包含head和tail
    curr := list.head
    idx := 0
    for {
        nodeIndex[curr] = idx
        if curr == list.tail {
            break
        } else {
            idx ++
            curr = curr.next[0]
        }
    }

    for i:=list.level-1; i>=0; i-- {
        curr := list.head
        for {
            fmt.Print(curr.key)
            if curr != list.head && curr != list.tail {
                repeat := width - len(strconv.Itoa(curr.key))
                if repeat > 0 {
                    fmt.Print(strings.Repeat(" ", repeat))
                }
            }

            if curr == list.tail {
                fmt.Println("")
                break
            } else {
                next := curr.next[i]
                diff := nodeIndex[next] - nodeIndex[curr] - 1
                gap := diff * (width + 1) + 1
                fmt.Print(strings.Repeat(" ", gap))
                curr = next
            }
        }
    }
}
