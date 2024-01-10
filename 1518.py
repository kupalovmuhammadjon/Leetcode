class Solution:
    def numWaterBottles(self, numBottles: int, numExchange: int) -> int:
        
        mx = numBottles
        empty = numBottles
        
        while empty >= numExchange:
            
            tmp = empty // numExchange
            mx += tmp
            empty = tmp + empty % numExchange

        
        return mx

print(Solution().numWaterBottles(15, 4))