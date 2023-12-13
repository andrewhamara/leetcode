using std::count;
class Solution {
  public:
    int numSpecial(vector<vector<int>>& mat) {
      int rows = mat.size();
      int cols = mat[0].size();

      vector<int> rowCount(rows);
      vector<int> colCount(cols);

      int result = 0;

      for (int i = 0; i < rows; i++) {
        for (int j = 0; j < cols; j++) {
          if (mat[i][j]) {
            rowCount[i]++;
            colCount[j]++;
          }
        }
      }

      for (int i = 0; i < rows; i++) {
        for (int j = 0; j < cols; j++) {
          if (mat[i][j] && rowCount[i] == 1 && colCount[j] == 1) {
            result++;
          }
        }
      }

      return result;
    }
};
