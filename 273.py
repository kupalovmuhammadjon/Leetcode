class Solution:
    
    def hundreds(self, lsnum):
        ones = {
            1: "One",
            2: "Two",
            3: "Three",
            4: "Four",
            5: "Five",
            6: "Six",
            7: "Seven",
            8: "Eight",
            9: "Nine",
        }

        tens = {
            1: "Ten",
            2: "Twenty",
            3: "Thirty",
            4: "Forty",
            5: "Fifty",
            6: "Sixty",
            7: "Seventy",
            8: "Eighty",
            9: "Ninety",
        }

        weirdos = {
            0: "Ten",
            1: "Eleven",
            2: "Twelve",
            3: "Thirteen",
            4: "Fourteen",
            5: "Fifteen",
            6: "Sixteen",
            7: "Seventeen",
            8: "Eighteen",
            9: "Nineteen",
        }
        
        h = "Hundred"
        if lsnum.count(0) == 3:
            return ""
        
        if len(lsnum) <= 3:
            if len(lsnum) == 1:
                return ones[lsnum[0]]
            elif len(lsnum) == 2:
                if lsnum[0] == 1:
                    return weirdos[lsnum[1]]
                elif lsnum[1] != 0:
                    return f"{tens[lsnum[0]]} {ones[lsnum[1]]}"
                else:
                    return tens[lsnum[0]]
            elif len(lsnum) == 3:
                res = f"{ones[lsnum[0]]} {h} "
                if lsnum[2] != 0:
                    res += f"{tens[lsnum[1]]} {ones[lsnum[2]]} "
                else:
                    res += tens[lsnum[1]]
                
                return res  
        else:
            return ""
            
    def numberToWords(self, num: int) -> str:
        if num == 0:
            return "Zero"

        lsnum = []
        if num <= 999:
            while num > 0:
                lsnum.append(num % 10)
                num //= 10
            lsnum.reverse()
            return self.hundreds(lsnum)

        else:
            t = "Thousand"
            m = "Million"
            b = "Billion"
            thousands = []
            temp = []
            while num > 0:
                if len(temp) == 3:
                    temp.reverse()
                    thousands.append(temp)
                    temp = []
                temp.append(num % 10)
                num //= 10
            temp.reverse()
            thousands.append(temp)
            temp = []
            thousands.reverse()
            
            for i in thousands:
                one = False
                two = False
                j = 0
                while j < len(i):
                    if i[j] == 0:
                        i.remove(0)
                        one = True
                    elif i[j] == 0 and one:
                        two = True
                        i.remove(0)
                    elif i[j] == 0 and one and two:
                        i.remove(0)
                    else:
                        j += 1
                    
            if len(thousands) == 2:
                return f"{self.hundreds(thousands[0])} {t} {self.hundreds(thousands[1])}"
            elif len(thousands) == 3:
                return f"{self.hundreds(thousands[0])} {m} {self.hundreds(thousands[1])} {t} {self.hundreds(thousands[2])}"
            elif len(thousands) == 4:
                return f"{self.hundreds(thousands[0])} {b} {self.hundreds(thousands[1])} {m} {self.hundreds(thousands[2])} {t} {self.hundreds(thousands[3])}"
                
                


# Test cases
print(Solution().numberToWords(0))
print(Solution().numberToWords(10))
print(Solution().numberToWords(20))
print(Solution().numberToWords(17))
print(Solution().numberToWords(450))
print(Solution().numberToWords(739))
print(Solution().numberToWords(999))
print(Solution().numberToWords(120091))
print(Solution().numberToWords(1000010))
print(Solution().numberToWords(1000010001))