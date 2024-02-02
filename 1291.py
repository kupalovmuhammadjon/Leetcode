class Solution:
    def sequentialDigits(self, low: int, high: int) -> List[int]:
        sc = '123456789'
        ls = []
        for i in range(len(sc)):
            for j in range(i+1, len(sc)+1):
                if low <= int(sc[i:j]) <= high:
                    ls.append(int(sc[i:j]))
                    
        return sorted(ls)
        