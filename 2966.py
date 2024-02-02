class Solution:
    def divideArray(self, nums: List[int], k: int) -> List[List[int]]:
        if len(nums) < 3:
            return []
        res = []
        nums.sort()
        
        temp = []
        for i in nums:
            if len(temp) == 3:
                res.append(temp)
                if max(temp) -  min(temp) > k:
                    return []
                temp = []
            temp.append(i)
        if len(temp) == 3:
                res.append(temp)
                if max(temp) -  min(temp) > k:
                    return []
                temp = []
        else:
            return []
        
        
        return res


        
        