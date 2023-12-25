class Solution:
    def findLucky(self, arr: List[int]) -> int:
        
        st = set(arr)
        mx = -1
        for i in st:
            if arr.count(i) == i and mx < i:
                mx = i      
            
        return mx