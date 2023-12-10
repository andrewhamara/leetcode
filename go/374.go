func guessNumber(n int) int {

    left := 1
    right := n

    for true {
        mid := (left + right) / 2

        result := guess(mid)

        switch result {
        case -1:
            right = mid - 1
        case 0:
            return mid
        case 1:
            left = mid + 1
        }
    }
    return -1
}
