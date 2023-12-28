class Solution:
    def finalPrices(self, prices: List[int]) -> List[int]:
        
        res = []
        for i in range(len(prices)):
            lamp = 1
            for j in range(i + 1, len(prices), 1):
                if prices[i] >= prices[j]:
                    res.append(prices[i] - prices[j])
                    lamp = 0
                    break
            if lamp:
                res.append(prices[i])
                
        return res
                
            