package normal

func ReverseString(s []byte)  {
    i := 0
    j := len(s) - 1

    for {
        if i >= j {
            break
        }

        tmp := s[i]
        s[i] = s[j]
        s[j] = tmp

        i ++
        j --
    }
}