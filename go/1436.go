func destCity(paths [][]string) string {
    s := make(map[string]bool)
    for _,value := range(paths) {
        s[value[0]] = true
    }
    for _,value := range(paths) {
        _,ok := s[value[1]]
        if !ok {
            return value[1]
        }
    }
    return ""
}
