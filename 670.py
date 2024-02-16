class Solution:
    def maximumSwap(self, num: int) -> int:
        ls = list(str(num))
        for i in range(len(ls)):
            if int(ls[i]) < int(max(ls[i:])):
                save = ls[i]
                ls[i] = max(ls[i:])
                j = len(ls)-1
                while j > i:
                    if ls[j] == ls[i]:
                        ls[j] = save
                        break
                    j -= 1
                break
        
        return int("".join(ls))

