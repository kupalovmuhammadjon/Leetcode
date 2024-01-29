class Solution:
    def setZeroes(self, matrix: List[List[int]]) -> None:
        l = len(matrix)
        rl = len(matrix[0])
        cor = []
        for i in range(l):
            for j in range(rl):
                if matrix[i][j] == 0:
                    cor.append([i, j])

        for i in cor:
            for k in range(rl):
                matrix[i[0]][k] = 0
            for k in range(l):
                matrix[k][i[1]] = 0
