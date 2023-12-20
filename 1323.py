class Solution:
    def maximum69Number (self, num: int) -> int:
        
        ls = list(str(num))
        
        for i in range(len(ls)):
            
            if ls[i] == "6":
                ls[i] = "9"
                break
        return int("".join(ls))