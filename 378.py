class Solution:
    def kthSmallest(self, matrix: List[List[int]], k: int) -> int:
        ls = []
        for i in range(len(matrix)):
            for j in range(len(matrix[0])):
                ls.append(matrix[i][j])

        ls.sort()
        return ls[k-1]