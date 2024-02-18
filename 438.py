class Solution:
    def findAnagrams(self, s: str, p: str) -> List[int]:
        res = []
        p =  sorted(p)
        for i in range(0, len(s)-len(p)+1, 1):
            st = sorted(s[i:i+len(p)])
            if st == p:
                res.append(i)
        
        return res


