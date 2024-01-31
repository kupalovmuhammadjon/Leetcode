class Solution:
    def findCenter(self, edges: List[List[int]]) -> int:
        n = edges[0][0]
        if n not in edges[1]:
            n = edges[0][1]
        return n