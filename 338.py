class Solution:
    def countBits(self, n: int) -> List[int]:
        ls = []
        for i in range(n+1):
            s = bin(i)
            ls.append(s.count("1"))
        return ls
    
    