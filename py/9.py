class Solution:
    def isPalindrome(self, x: int) -> bool:
        if x < 0: return False
        temp,rev = x, 0
        while temp:
            rev = rev * 10 + (temp % 10)
            temp //= 10
        return rev == x
