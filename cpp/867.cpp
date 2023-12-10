class Solution {
  public:
    vector<vector<int>> transpose(vector<vector<int>>& matrix) {
      int rowCount = matrix.size();
      int colCount = matrix[0].size();
      std::vector<std::vector<int>> transposed(colCount);

      for (int i = 0; i < colCount; i++)
        for (int k = 0; k < rowCount; k++)
          transposed[i].push_back(matrix[k][i]);

      return transposed;
    }
};
