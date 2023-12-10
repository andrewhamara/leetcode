func firstUniqChar(s string) int {
    m := make(map[rune]int)

    for _, chr := range s {
        _, ok := m[chr]
        if ok {
            m[chr] += 1
        }else {
            m[chr] = 1
        }
    }

    for i,c := range s{
        if m[c] < 2{
            return i
        }
    }
    return -1
}
