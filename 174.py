from itertools import product
from typing import List
class Solution:
    def calculateMinimumHP(self, dungeon: List[List[int]]) -> int:
        oper = ["r", "d"]
        comb = product(oper, repeat=len(dungeon) + len(dungeon[0])-2)

        mx = 0
        for c in comb:
            i = 0
            j = 0
            sm = 0
            if c.count("r") > len(dungeon[0])-1 or c.count("d") > len(dungeon[0])-1:
                continue
            print(c)
            
dungeon = [[-2,-3,3],[-5,-10,1],[10,30,-5]]
print(Solution().calculateMinimumHP(dungeon))