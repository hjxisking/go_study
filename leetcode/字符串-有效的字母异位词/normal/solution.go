package normal

func IsAnagram(s string, t string) bool {
    hashMap := map[rune]int{}
    for _, chr := range []rune(s) {
        if _, ok := hashMap[chr]; ok {
            hashMap[chr] += 1
        } else {
            hashMap[chr] = 1
        }
    }

    for _, chr := range []rune(t) {
        if _, ok := hashMap[chr]; ok {
            hashMap[chr] -= 1
            if hashMap[chr] == 0 {
                delete(hashMap, chr)
            }
        } else {
            return false
        }
    }

    return len(hashMap) == 0
}