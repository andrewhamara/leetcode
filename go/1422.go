func maxScore(s string) int {
    onesToRight := 0
    for _,c := range s {
      if c == '1' {
        onesToRight++
      }
    }
    zerosToLeft, maxScore := 0,0

    for i:=0; i < len(s)-1; i++ {
        if s[i] == '0' {
          zerosToLeft++
        } else {
          onesToRight--
        }
        if maxScore < (onesToRight + zerosToLeft) {
          maxScore = onesToRight + zerosToLeft
        }
    }
    return maxScore
}
