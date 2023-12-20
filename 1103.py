from typing import List

class Solution:
    def distributeCandies(self, candies: int, num_people: int) -> List[int]:
        d = 1
        res = [0] * num_people
        i = 0

        while candies > 0:
            if candies >= d:
                res[i] += d
                candies -= d
                d += 1
            else:
                res[i] += candies
                candies = 0

            if i == num_people - 1:
                i = 0
            else:
                i += 1

        return res

# Done passed from all tests