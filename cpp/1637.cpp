class Solution {
public:
  int maxWidthOfVerticalArea(vector<vector<int>>& points) {
    vector<int> xVals;
    for (auto &point : points)
      xVals.push_back(point[0]);

    sort(xVals.begin(), xVals.end());

    int diff = 0;

    for (int i = 1; i < xVals.size(); i++) {
      int cur = xVals[i] - xVals[i-1];
      if (cur > diff)
        diff = cur;
    }
    return diff;
  }
};
