class Solution:
    def maxWidthOfVerticalArea(self, points: List[List[int]]) -> int:
        xVals = sorted([x[0] for x in points])
        diff = -100000

        for i in range(1, len(xVals)):
            diff = max(diff, xVals[i] - xVals[i-1])
        return diff
