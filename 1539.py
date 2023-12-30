class Solution:
    def findKthPositive(self, arr: List[int], k: int) -> int:
        
        c = 0
        i = 1
        while True:
            if i not in arr:
                c += 1
            if c == k:
                return i
            i += 1
            
        return -1