class Solution:
    def replaceElements(self, arr: List[int]) -> List[int]:
        n = len(arr)
        res = [-1] 
        
        max_right = -1  
        
        for i in range(n - 1, 0, -1):
            max_right = max(max_right, arr[i]) 
            res.insert(0, max_right) 
            
        return res