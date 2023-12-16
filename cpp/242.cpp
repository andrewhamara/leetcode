class Solution {
public:
  bool isAnagram(string s, string t) {
    vector<int> vals(26,0);
  for (int i = 0; i < s.length(); i++)
      vals[s[i] - 'a']++;
  for (int k = 0; k < t.length(); k++)
      vals[t[k] - 'a']--;
    for (int x : vals)
      if (x != 0) {return false;}
    return true;
  }
};
