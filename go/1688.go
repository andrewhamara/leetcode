func numberOfMatches(n int) int {
    matches := 0
    for n > 1 {
        if n % 2 == 0 {
            n /= 2
            matches += n
        } else {
            n = ((n-1) / 2) + 1
            matches += n - 1
        }
    }
    return matches
}
