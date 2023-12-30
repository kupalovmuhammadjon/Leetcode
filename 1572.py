class Solution:
    def diagonalSum(self, mat: List[List[int]]) -> int:
        
        m = len(mat)
        sm = 0
        for i in range(m):
            for j in range(m):
                if i == j or i + j == m - 1:
                    sm += mat[i][j]
                    
        return sm
                