class Solution:
    def isAnagram(self, s:str, t:str) -> bool:
        vals = [0] * 26
        for x in s:
            vals[ord(x) - 97] += 1
        for y in t:
            vals[ord(y) - 97] -= 1
        return all(x == 0 for x in vals)
