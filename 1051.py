class Solution:
    def heightChecker(self, heights: List[int]) -> int:
        
        ls = sorted(heights)
        df = 0
        
        for i in range(len(ls)):
            if heights[i] != ls[i]:
                df += 1
        return df
    
# Done Passed from all tests
