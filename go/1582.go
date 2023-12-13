func numSpecial(mat [][]int) int {

  rows := len(mat)
  cols := len(mat[0])

  rowCount := make([]int, rows)
  colCount := make([]int, cols)

  result := 0

  for i:=0; i < rows; i++ {
      for k:=0; k < cols; k++ {
          if mat[i][k] == 1 {
              rowCount[i] += 1
              colCount[k] += 1
          }
      }
  }

  for i:=0; i < rows; i++ {
      for k:=0; k < cols; k++ {
          if mat[i][k] == 1 && rowCount[i] == 1 && colCount[k] == 1 {
              result += 1
          }
      }
  }

  return result
}
