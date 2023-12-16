class Solution:
    def removeDuplicates(self, nums: List[int]) -> int:
        i = 0
        for x in nums:
            if 2 > i or nums[i-2] < x:
                nums[i] = x
                i += 1
        return i
