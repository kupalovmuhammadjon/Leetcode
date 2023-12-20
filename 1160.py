from typing import List

class Solution:
    def countCharacters(self, words: List[str], chars: str) -> int:
        org = list(chars)
        lens = 0

        for i in words:
            ls = list(org)
            c = 1

            for j in i:
                if j in ls:
                    ls.remove(j)
                else:
                    c = 0
                    break

            if c:
                lens += len(i)
                org = list(chars) 

        return lens
    
# Done passed from all tests

