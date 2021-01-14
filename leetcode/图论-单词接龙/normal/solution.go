package normal

func LadderLength(beginWord string, endWord string, wordList []string) int {
    edge := map[string][]string{}
    for _, word := range wordList {
        addEdge(word, edge)
    }
    if _, ok := edge[endWord]; !ok {
        return 0
    }

    addEdge(beginWord, edge)

    queue := []string{beginWord}
    dist := map[string]int{}
    dist[beginWord] = 1
    for len(queue) > 0 {
        current := queue[0]
        queue = queue[1:]
        for _, word := range edge[current] {
            if word == endWord {
                return dist[current] / 2 + 1
            } else {
                if _, ok := dist[word]; !ok {
                    dist[word] = dist[current] + 1
                    queue = append(queue, word)
                }
            }
        }
    }
    return 0
}

func addEdge(word string, edge map[string][]string) {
    if _, ok := edge[word]; !ok {
        edge[word] = make([]string, 0)
        s := []byte(word)
        // hot -> *ot, h*t, ho*， 增加虚拟节点，如果另一个word的虚拟节点切好右一个和hot的虚拟节点一样，就说明这2个word可以联通（只差一个字符）
        for i, char := range s {
            s[i] = '*'
            tmp := string(s)
            edge[word] = append(edge[word], tmp)
            if _, ok := edge[tmp]; !ok {
                edge[tmp] = make([]string, 0)
            }
            edge[tmp] = append(edge[tmp], word)
            s[i] = char
        }
    }
}