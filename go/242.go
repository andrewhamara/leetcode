func isAnagram(s string, t string) bool {
  checker := make([]int, 26)
  for i:=0; i < len(s); i++ {
    checker[s[i] - 'a']++
  }
  for k:=0; k < len(t); k++ {
    checker[t[k] - 'a']--
  }

  for _,val := range checker {
    if val != 0 {
      return false
    }
  }
  return true
}
