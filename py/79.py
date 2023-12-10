# Author: Andrew Hamara

# Solution to leetcode problem 79. Word Search

class Solution:
    def exist(self, board: List[List[str]], word:str) -> bool:

        rowCount = len(board)
        colCount = len(board[0]) # Or any index that is guaranteed to exist...

        path = set()
        wordLen = len(word)

        def dfs(row, col, idx):
            if idx == wordLen
                return True

            if (row < 0 or
                col < 0 or
                row >= rowCount or
                col >= colCount or
                [row, col] in path or
                board[row][col] != word[idx]):
                    return False

            path.add([row, col])
            valid = (dfs(row + 1, col, i + 1) or
                     dfs(row - 1, col, i + 1) or
                     dfs(row, col + 1, i + 1) or
                     dfs(row, col - 1, i + 1))
            path.remove([row, col])
            return res

        for
