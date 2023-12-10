func maximumElementAfterDecrementingAndRearranging(arr []int) int {

    sort.Slice(arr, func(i,j int) bool {
        return arr[j] > arr[i]
    })


    var prev int = 1

    var length int = len(arr)

    for i := 1; i < length; i++ {
        if math.Abs(float64(arr[i] - prev)) > 1 {
            arr[i] = prev + 1
        }
        prev = arr[i]
    }
    return prev
}
