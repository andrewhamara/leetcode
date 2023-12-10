# Author: Andrew Hamara

# Solution for leetcode problem 231. Power of Two

class Solution:
    def isPowerOfTwo(self, n:int) -> bool:
        if n <= 0:
            return False

        while n % 2 == 0:
            n /= 2

        return n == 1
