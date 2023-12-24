func minOperations(s string) int {
    first := s[0]
    second := '0'
    if first == '0' {
        second = '1'
    } else {
        second = '0'
    }
    result := 0
    for i,c := range(s) {
        if i % 2 == 0 && c != rune(first) {
            result++
        } else if i % 2 == 1 && c != rune(second) {
            result++
        }
    }
    opposite := len(s) - result
    if result > opposite {
        return opposite
    }
    return result
}
