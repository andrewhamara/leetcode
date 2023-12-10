func transpose(matrix [][]int) [][]int {
    rowCount := len(matrix)
    colCount := len(matrix[0])
    transposed := make([][]int, colCount)

    for i:=0; i < colCount; i++ {
        for k:=0; k < rowCount; k++ {
            transposed[i] = append(transposed[i], matrix[k][i])
        }
    }

    return transposed
}
