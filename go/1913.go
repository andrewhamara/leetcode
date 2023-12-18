func maxProductDifference(nums []int) int {
    max1, max2 := -1000000, -1000000
    min1, min2 := 1000000, 1000000

    for _,num := range nums {
        if num <= min1 {
            min2 = min1
            min1 = num
        } else if num < min2 {
            min2 = num
        }
        if num >= max1 {
            max2 = max1
            max1 = num
        } else if num > max2 {
            max2 = num
        }
    }
    return (max1 * max2) - (min2 * min1)
}
