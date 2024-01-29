class Solution:
    
    def hundreds(self, lsnum):
        ones = {
            0: "",
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
            0: "",
            1: "Eleven",
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
        res = ""
        if len(lsnum) <= 3:
            if len(lsnum) == 1:
                res += ones[lsnum[0]]
            elif len(lsnum) == 2:
                if lsnum[0] == 1:
                    res += weirdos[lsnum[1]]
                elif lsnum[1] != 0:
                    res += f"{tens[lsnum[0]]} {ones[lsnum[1]]}"
                else:
                    res += tens[lsnum[0]]
            elif len(lsnum) == 3:
                res = f"{ones[lsnum[0]]} {h} "
                if lsnum[2] != 0 and lsnum[1] != 1:
                    res += f"{tens[lsnum[1]]} {ones[lsnum[2]]}"
                elif lsnum[2] != 0 and lsnum[1] == 1:
                    if lsnum[1] != 0:
                        res += f"{tens[lsnum[1]]}"
                else:
                    if lsnum[-1] == 0:
                        res += f"{weirdos[0]}"
                    else:
                        res += tens[lsnum[1]]
                    
            res =  list(res.strip())
            i = 1
            while i < len(res):
                if res[i] == " " and res[i-1] == " ":
                    res.pop(i)
                else:
                    i += 1   

            return "".join(res)
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
                while i and i[-1] == 0:
                    i.pop()
                if not i:
                    i.append(0)


                    
            res = ""
            if len(thousands) == 2:
                n = self.hundreds(thousands[0])
                if len(n) != 0:
                    res += f"{n} {t} "
                n = self.hundreds(thousands[1])
                if len(n) != 0:
                    res += n
            
            elif len(thousands) == 3:
                n = self.hundreds(thousands[0])
                if len(n) != 0:
                    res += f"{n} {m} "
                n = self.hundreds(thousands[1])
                if len(n) != 0:
                    res += f"{n} {t} "
                n = self.hundreds(thousands[2])
                if len(n) != 0:
                    res += n
            
            elif len(thousands) == 4:
                n = self.hundreds(thousands[0])
                if len(n) != 0:
                    res += f"{n} {b} "
                    
                n = self.hundreds(thousands[1])
                if len(n) != 0:
                    res += f"{n} {m} "
                n = self.hundreds(thousands[2])
                if len(n) != 0:
                    res += f"{n} {t} "
                n = self.hundreds(thousands[3])
                if len(n) != 0:
                    res += n
               
            res =  list(res.strip())
            
            i = 1
            while i < len(res):
                if res[i] == " " and res[i-1] == " ":
                    res.pop(i)
                else:
                    i += 1 

            return "".join(res)
                
                


# Test cases

print(Solution().numberToWords(101))
print(Solution().numberToWords(111))
print(Solution().numberToWords(110))
print(Solution().numberToWords(739))
print(Solution().numberToWords(999))
print(Solution().numberToWords(120091))
print(Solution().numberToWords(1000010))
print(Solution().numberToWords(1000010001))
