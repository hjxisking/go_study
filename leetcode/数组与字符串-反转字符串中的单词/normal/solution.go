package normal

func ReverseWords(s string) string {
    s1 := []byte(s)
    s1 = append(s1, ' ')
    i,j := 0,0
    for idx, b := range s1 {
        if b == ' ' {
            j = idx - 1
            for i <= j {
                tmp := s1[i]
                s1[i] = s1[j]
                s1[j] = tmp

                i ++
                j --
            }
            i = idx + 1
        }
    }

    return string(s1[0:len(s)])
}