class Solution:
    def canMakeArithmeticProgression(self, arr: List[int]) -> bool:
        arr.sort()
        
        d = arr[1] - arr[0]
        
        for i in range(len(arr)-1):
            if arr[i] + d != arr[i+1]:
                return False
        
        return True