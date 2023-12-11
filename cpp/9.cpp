class Solution {
  public:
    bool isPalindrome(int x) {
      if (x < 0) return false;
      long long temp = x, reversed = 0;
      while (temp) {
        reversed = reversed * 10 + (temp % 10);
        temp /= 10;
      }
      return reversed == x;
    }
};
