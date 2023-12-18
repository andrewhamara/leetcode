class Solution {
public:
  int maxProductDifference(vector<int>& nums) {
    int max1 = -1000000,  max2 = -1000000;
    int min1 =  1000000,  min2  = 1000000;

    for (int n : nums) {
      if (n <= min1) {
        min2 = min1;
        min1 = n;
      }
      else if (n < min2) {
        min2 = n;
      }
      if (n >= max1) {
        max2 = max1;
        max1 = n;
      }
      else if (n > max2) {
        max2 = n;
      }
    }
    return (max1 * max2) - (min2 * min1);
  }
};
