class Solution:
    def maxScore(self, s: str) -> int:
        mx = 0
        for i in range(1, len(s), 1):
           zero = list(s[:i])
           one = list(s[i:])
           if mx < zero.count("0") + one.count("1"):
               mx = zero.count("0") + one.count("1")
               
        return mx
               