class Solution:
    def runningSum(self, nums: List[int]) -> List[int]:
        
        sumLs = []
        runningSum = 0
        for i in nums:
            runningSum += i
            sumLs.append(runningSum)
            
        return sumLs