class Solution:
    def onesMinusZeros(self, grid: List[List[int]]) -> List[List[int]]:
        onesRow = [0 for _ in range(len(grid))]
        zerosRow = [0 for _ in range(len(grid))]
        onesCol = [0 for _ in range(len(grid[0]))]
        zerosCol = [0 for _ in range(len(grid[0]))]

        for i in range(len(grid)):
            for k in range(len(grid[0])):
                if grid[i][k] == 0:
                    zerosRow[i] += 1
                    zerosCol[k] += 1
                else:
                    onesRow[i] += 1
                    onesCol[k] += 1

        diff = [[] for _ in range(len(grid))]
        for i in range(len(grid)):
            for k in range(len(grid[0])):
                diff[i].append(onesRow[i] + onesCol[k] - zerosRow[i] - zerosCol[k])
        return diff
