package normal

func LongestCommonPrefix(strs []string) string {
    length := len(strs)
    if length == 0 {
        return ""
    }

    if length == 1{
        return strs[0]
    }

    commonPrefix := []byte{}
    i := 0

    for {
        var tmp byte
        for _, str := range strs {
            if i >= len(str) {
                return string(commonPrefix)
            } else {
                if tmp == 0 {
                    tmp = str[i]
                } else {
                    if tmp != str[i] {
                        return string(commonPrefix)
                    }
                }
            }
        }
        i ++
        commonPrefix = append(commonPrefix, tmp)
    }
}