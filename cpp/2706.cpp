class Solution {
public:
  int buyChoco(vector<int>& prices, int money) {
    int min1 = 128, min2 = 128;
    for (int n : prices) {
      if (min1 >= n) {
        min2 = min1;
        min1 = n;
      }
      else if (min2 > n) {
        min2 = n;
      }
    }
    int leftover = money - min1 - min2;
    return leftover >= 0 ? leftover : money;
  }
};
