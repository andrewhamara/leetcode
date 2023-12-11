class Solution {
  public:
    int findSpecialInteger(vector<int>& arr) {
      float qtr = arr.size() / 4.0;
      std::map<int,int> freq;
      for (int n : arr ) {
        freq[n]++;
        int cur = freq[n];
        if (float(cur) > qtr)
          return n;
      }
      return -1;
    }
};
