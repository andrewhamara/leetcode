func countBits(n int) []int {
    length := n + 1
    bits :=make([]int, length)
    offset := 1

    for i := 1; i < length; i++ {
        if i == offset * 2 {
            offset = i
        }
        bits[i] = 1 + bits[i-offset]
    }
    return bits
}
