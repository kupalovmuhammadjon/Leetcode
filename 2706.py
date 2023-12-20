class Solution:
    def buyChoco(self, prices: List[int], money: int) -> int:
        c1 = min(prices)
        prices.remove(c1)
        c2 = min(prices)
        if money - (c1 + c2) >= 0:
            return money - (c1 + c2)
        else:
            return money

            
# Done