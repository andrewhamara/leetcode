func onesMinusZeros(grid [][]int) [][]int {
  rows, cols := len(grid), len(grid[0])
  onesRow  := make([]int, rows)
  zerosRow := make([]int, rows)
  onesCol  := make([]int, cols)
  zerosCol := make([]int, cols)

  for i:=0; i < rows; i++ {
    for j:=0; j < cols; j++ {
      if grid[i][j] == 1 {
        onesRow[i]++
        onesCol[j]++
      } else {
        zerosRow[i]++
        zerosCol[j]++
      }
    }
  }

  diff := make([][]int, rows)
  for i:=0; i < rows; i++ {
    for j:=0; j < cols; j++ {
      diff[i] = append(diff[i], onesRow[i] + onesCol[j] - zerosRow[i] - zerosCol[j])
    }
  }
  return diff
}
