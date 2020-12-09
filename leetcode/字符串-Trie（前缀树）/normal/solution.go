package normal

/**
 * Your Trie object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Insert(word);
 * param_2 := obj.Search(word);
 * param_3 := obj.StartsWith(prefix);
 */

type Trie struct {
    root *node
}

type node struct {
    value byte
    next map[byte]*node
    isEnd bool
}

/** Initialize your data structure here. */
func Constructor() Trie {
    return Trie{root:&node{
        value: 0,
        next:  map[byte]*node{},
    }}
}


/** Inserts a word into the trie. */
func (this *Trie) Insert(word string)  {
    current := this.root
    for _, char := range []byte(word) {
        if next, ok := current.next[char]; !ok {
            newNode := &node{
                value: char,
                next:  map[byte]*node{},
            }
            current.next[char] = newNode
            current = newNode
        } else {
            current = next
        }
    }
    current.isEnd = true
}


/** Returns if the word is in the trie. */
func (this *Trie) Search(word string) bool {
    current := this.root
    for _, char := range []byte(word) {
        if next, ok := current.next[char]; ok {
            current = next
        } else {
            return false
        }
    }

    return current.isEnd
}


/** Returns if there is any word in the trie that starts with the given prefix. */
func (this *Trie) StartsWith(prefix string) bool {
    current := this.root
    for _, char := range []byte(prefix) {
        if next, ok := current.next[char]; ok {
            current = next
        } else {
            return false
        }
    }

    return true
}


