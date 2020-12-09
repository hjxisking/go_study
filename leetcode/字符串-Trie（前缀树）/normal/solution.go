package normal

/**
 * Your Trie object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Insert(word);
 * param_2 := obj.Search(word);
 * param_3 := obj.StartsWith(prefix);
 */

type Trie struct {
    value byte
    next map[byte]*Trie
    isEnd bool
}

/** Initialize your data structure here. */
func Constructor() Trie {
    return Trie{
        value: 0,
        next:  map[byte]*Trie{},
        isEnd: false,
    }
}


/** Inserts a word into the trie. */
func (this *Trie) Insert(word string)  {
    current := this
    for _, char := range []byte(word) {
        if next, ok := current.next[char]; !ok {
            node := &Trie{
                value: char,
                next:  map[byte]*Trie{},
                isEnd: false,
            }

            current.next[char] = node
            current = node
        } else {
            current = next
        }
    }
    current.isEnd = true
}


/** Returns if the word is in the trie. */
func (this *Trie) Search(word string) bool {
    current := this
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
    current := this
    for _, char := range []byte(prefix) {
        if next, ok := current.next[char]; ok {
            current = next
        } else {
            return false
        }
    }

    return true
}

func (this *Trie) ContainKey(k byte) (Trie,bool) {
    if node, ok := this.next[k]; ok {
        return *node, true
    } else {
        return Trie{}, false
    }
}
