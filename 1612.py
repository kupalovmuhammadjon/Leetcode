from typing import List
class Solution:
    def decrypt(self, code: List[int], k: int) -> List[int]:
        
        res = []
        if k > 0:
            for i in range(len(code)):
                c = 0
                sm = 0
                j = i+1
                while True:
                    if j == len(code):
                        j = 0
                    sm += code[j]
                    c += 1
                    j = j + 1
                    if c == k:
                        break
                    
                res.append(sm)
        elif k < 0:
            
            for i in range(len(code)):
                c = 0
                sm = 0
                j = i-1
                while True:
                    if j < 0:
                        j = len(code)-1

                    c += 1
                    sm += code[j]
                    
                    j -= 1
                    if j == len(code):
                        j = 0
                    if c == abs(k):
                        break
                res.append(sm)
        else:
            for i in range(len(code)):
                res.append(0)
                
        return res
            