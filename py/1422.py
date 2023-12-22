class Solution:
    def maxScore(self, s: str) -> int:
        onesToRight = s.count('1')
        zerosToLeft = 0
        maxScore = 0
        for i in range(len(s) - 1):
            x = s[i]
            if x == '0':
                zerosToLeft += 1
            else:
                onesToRight -= 1
            maxScore = max(maxScore, zerosToLeft + onesToRight)
        return maxScore
