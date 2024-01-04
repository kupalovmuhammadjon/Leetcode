class Solution:
    def maximumWealth(self, accounts: List[List[int]]) -> int:

        bigBoy = 0
        for i in accounts:
            if sum(i) > bigBoy:
                bigBoy = sum(i)
                
        return bigBoy