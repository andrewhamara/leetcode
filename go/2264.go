func largestGoodInteger(num string) string {
    best := ""
    length := len(num)

    for i := 0; i < length - 2; i++ {
        if (num[i] == num[i+1] && num[i+1] == num[i+2]) {
            cur := num[i]
            if (best == "" || best[0] < cur) {
                best = string(cur) + string(cur) + string(cur)
            }
        }
    }
    return best
}
