class Solution:
    def commonChars(self, words: List[str]) -> List[str]:
        
        res = []
        
        for i in words[0]:
            c = 1
            for j in words:
                if i not in j:
                    c = 0
            if c:
                res.append(i)
        return res
    
# Done passed from all tests