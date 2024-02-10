class Solution:
    def isAdditiveNumber(self, num: str) -> bool:
        for i in range(1, len(num)):
            for j in range(i + 1, len(num)):
                f1 = int(num[:i])
                f2 = int(num[i:j])
                if (len(str(f1)) != i or len(str(f2)) != j - i):
                    break
                if self.isValid(num[j:], f1, f2):
                    return True
        return False
    
    def isValid(self, num, f1, f2):
        while num:
            f3 = f1 + f2
            if num.startswith(str(f3)):
                f1, f2 = f2, f3
                num = num[len(str(f3):]
            else:
                return False
        return True
