class Solution:
    def destCity(self, paths: List[List[str]]) -> str:
        atZero = set()
        allVals = set()
        for path in paths:
            atZero.add(path[0])
            allVals.add(path[0])
            allVals.add(path[1])
        return (allVals ^ atZero).pop()
