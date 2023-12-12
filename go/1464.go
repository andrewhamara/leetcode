func maxProduct(nums []int) int {
  first, second := -1, -1
  for _, num := range nums {
    if num >= first {
      first, second = num, first
    } else if num > second {
      second = num}
    }
  return (first - 1) * (second - 1)
}
