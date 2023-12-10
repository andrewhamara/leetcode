func hammingDistance(x int, y int) int {
    diff := 0

    for i:=0; i < 32; i++ {
        diff += (x%2) ^ (y%2)
        x >>=1
        y >>=1
    }
    return diff
}
