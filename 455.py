class Solution:
    def findContentChildren(self, g: List[int], s: List[int]) -> int:
        
        g.sort()
        s.sort()

        assigned = 0

        for greed in g:
            for i in range(len(s)):
                if greed <= s[i]:
                    assigned += 1
                    s.pop(i)
                    break

        return assigned