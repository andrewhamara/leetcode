class Solution:
    def minOperations(self, s: str) -> int:
        res = 0
        first = s[0]
        second = '1' if first == '0' else '0'
        for i,c in enumerate(s):
            if i % 2 == 0 and c != first:
                res += 1
            elif i % 2 == 1 and c != second:
                res += 1
        return min(res, len(s) - res)
