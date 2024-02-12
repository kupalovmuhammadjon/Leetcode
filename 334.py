class Solution:
    def increasingTriplet(self, nums: List[int]) -> bool:
        mx = float("inf")
        mn = float("inf")
        for i in nums:

            if i <= mn:
                mn = i
            elif i <= mx:
                mx = i
            else:
                return 1
            
        return 0
