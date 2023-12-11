func findSpecialInteger(arr []int) int {
  qtr := len(arr) / 4.0
  freq := make(map[int]int)

  for _,v := range(arr) {
    freq[v]++
    if freq[v] > qtr {
      return v
    }
  }
  return -1
}
