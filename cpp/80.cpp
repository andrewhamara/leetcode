class Solution {
public:
  int removeDuplicates(vector<int>& nums) {
    int i = 0;
    for (int x : nums)
      if (2 > i || nums[i-2] < x)
        nums[i++] = x;
    return i;
  }
};
