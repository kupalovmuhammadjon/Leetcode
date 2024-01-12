class Solution:
    def maxDepth(self, s: str) -> int:
        ls = []
                
        count = 0
        for i in s:
            if i == "(":
                count += 1
                
            elif i == ")":
                count -= 1
            ls.append(count)
                
        return max(ls)
       