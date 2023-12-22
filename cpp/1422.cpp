class Solution {
  public:
    int maxScore(string s) {
      int onesToRight = 0;
      for (auto &c : s)
        if (c == '1')
          onesToRight++;
      int maxScore = 0, zerosToLeft = 0;

      for (int i = 0; i < s.length()-1; i++) {
        auto x = s[i];
        if (x == '0')
          zerosToLeft++;
        else
          onesToRight--;
        if (onesToRight + zerosToLeft > maxScore)
          maxScore = onesToRight + zerosToLeft;
      }
      return maxScore;
    }
};
