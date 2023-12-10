func totalMoney(n int) int {

  nums := make([]int, 1)
  nums[0] = 1

  for i:=1; i < n; i++ {
    if i % 7 == 0 {
      nums = append(nums, nums[i-7] + 1)
    } else {
      nums = append(nums, nums[i-1] + 1)
    }
  }

  result := 0
  for _,v := range nums {
    result += v
  }
  return result
}
