class Solution {
  public:
    int findSpecialInteger(vector<int>& arr) {
      float qtr = arr.size() / 4.0;
      std::map<int,int> freq;
      for (int n : arr ) {
        freq[n]++;
        if (float(freq[n]) > qtr)
          return n;
      }
      return -1;
    }
};
