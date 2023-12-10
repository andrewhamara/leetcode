func distributeCandies(candyType []int) int {
    set := make(map[int]bool)
    for _,candy := range candyType {
        set[candy] = true
    }

    unqLen := len(set)
    ogLen := len(candyType)
    r := 0

    for r < unqLen && r != ogLen/2 {
        r++
    }

    return r
}
