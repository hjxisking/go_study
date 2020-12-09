package normal

type Trie struct {
    value byte
    next map[byte]*Trie
    isEnd bool
}

/** Initialize your data structure here. */
func NewTrie() Trie {
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

func (this *Trie) ContainKey(k byte) (Trie,bool) {
    if node, ok := this.next[k]; ok {
        return *node, true
    } else {
        return Trie{}, false
    }
}

func (this *Trie) DeleteKey(k byte) {
    if _, ok := this.next[k]; ok {
        delete(this.next, k)
    }
}


func FindWords(board [][]byte, words []string) []string {
    trie := NewTrie()
    for _, word := range words {
        trie.Insert(word)
    }

    ans := []string{}
    for i := 0; i < len(board); i ++ {
        for j := 0; j < len(board[0]); j ++ {
            ans = append(ans, backtracking(board, i, j, trie)...)
        }
    }

    return ans
}

func backtracking(board [][]byte, x, y int, trie Trie) []string {
    res := []string{}
    prefix := board[x][y]
    board[x][y] = '#'       // 防止回溯时找到起始点变成环形，改变起始点为特殊字符，在结束时还原，当递归到起始点的时候必然不符合要找的单词而剪枝

    if next, ok := trie.ContainKey(prefix); ok {
        row := []int{-1 ,0, 1, 0}
        col := []int{0, 1, 0, -1}

        tmp := []string{}

        for i := 0; i < 4; i ++ {
            x1 := x + row[i]
            if x1 < 0 || x1 >= len(board) {
                continue
            }
            y1 := y + col[i]
            if y1 < 0 || y1 >= len(board[0]) {
                continue
            }

            tmp = append(tmp, backtracking(board, x1, y1, next)...)
        }

        if next.isEnd {
            res = []string{string(prefix)}
        }

        if len(tmp) > 0 {
            res = append(res, combine(prefix, tmp)...)
        }

        // 剪枝，可减少重复工作量以及防止结果重复
        if next.isEnd {
            trie.DeleteKey(prefix)
        }
    }

    board[x][y] = prefix
    return res
}

func combine(s byte, words []string) []string {
    res := []string{}
    for _, word := range words {
        res = append(res, string(s)+word)
    }
    return res
}