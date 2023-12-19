class Solution:
    def imageSmoother(self, img: List[List[int]]) -> List[List[int]]:
        smoothed = [[] for _ in range(len(img))]
        for i in range(len(img)):
            for k in range(len(img[0])):
                sumToPut = img[i][k]
                validSurrounding = 1
                upValid, downValid, rightValid, leftValid = False,False,False,False

                # valid straight to right
                if k+1 < len(img[0]):
                    sumToPut += img[i][k+1]
                    validSurrounding += 1
                    rightValid = True
                # valid straight to left
                if k-1 >= 0:
                    sumToPut += img[i][k-1]
                    validSurrounding += 1
                    leftValid = True
                # valid straight down
                if i+1 < len(img):
                    sumToPut += img[i+1][k]
                    validSurrounding += 1
                    downValid = True
                # valid straight up
                if i-1 >= 0:
                    sumToPut += img[i-1][k]
                    validSurrounding += 1
                    upValid = True
                # up to right
                if upValid and rightValid:
                    sumToPut += img[i-1][k+1]
                    validSurrounding += 1
                # up to left
                if upValid and leftValid:
                    sumToPut += img[i-1][k-1]
                    validSurrounding += 1
                # down to right
                if downValid and rightValid:
                    sumToPut += img[i+1][k+1]
                    validSurrounding += 1
                # down to left
                if downValid and leftValid:
                    sumToPut += img[i+1][k-1]
                    validSurrounding += 1
                smoothed[i].append(int(sumToPut / validSurrounding))
        return smoothed
