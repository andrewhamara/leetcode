class Solution:
    def numSpecial(self, mat: List[List[int]]) -> int:
        total = 0
        for i in range(len(mat)):
            row = mat[i]
            if row.count(1) == 1:
                colIdx = row.index(1)
                col = [r[colIdx] for r in mat]
                if col.count(1) == 1:
                    total += 1
        return total
