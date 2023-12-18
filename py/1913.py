class Solution:
    def maxProductDifference(self, nums: List[int]) -> int:
        min1, min2 = 1000000, 1000000
        max1, max2 = -1000000, -1000000

        for n in nums:
            if min1 >= n:
                min2 = min1
                min1 = n
            elif min2 > n:
                min2 = n
            if max1 <= n:
                max2 = max1
                max1 = n
            elif max2 < n:
                max2 = n
        return (max1 * max2) - (min2 * min1)
