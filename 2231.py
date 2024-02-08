import heapq
class Solution:
    def largestInteger(self, num: int) -> int:
        odds = []
        evens = []
        
        for i in str(num):
            if int(i) % 2:
                odds.append(-int(i))
            else:
                evens.append(-int(i))

        res = ""
        heapq.heapify(odds)
        heapq.heapify(evens)

        for i in str(num):
            n = -1 * int(i)
            if n % 2:
                mn = heapq.heappop(odds)
                res += str(-1 * mn)
            else:
                mn = heapq.heappop(evens)
                res += str(-1 * mn)
        
        return int(res)
