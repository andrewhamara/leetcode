func buyChoco(prices []int, money int) int {
  min1, min2 := 128,128
  for _,n := range prices {
    if min1 >= n {
      min2 = min1
      min1 = n
    } else if min2 > n {
      min2 = n
    }
  }
  leftover := money - min1 - min2
  if leftover >= 0 {
    return leftover
  }
  return money
}
