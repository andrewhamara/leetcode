class Solution:
    def buyChoco(self, prices:List[int], money:int) -> int:
        min1, min2 = 128,128
        for x in prices:
            if min1 >= x:
                min2 = min1
                min1 = x
            elif min2 > x:
                min2 = x
        leftover = money - min1 - min2
        return leftover if leftover >= 0 else money
