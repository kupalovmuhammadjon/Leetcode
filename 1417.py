class Solution:
    def reformat(self, s: str) -> str:
        digits = []
        symb = []
        res = ""

        for i in s:
            if i.isdigit():
                digits.append(i)
            else:
                symb.append(i)

        if abs(len(digits) - len(symb)) > 1:
            return ""

        if len(digits) < len(symb):
            digits, symb = symb, digits

        for i in range(len(symb)):
            res += digits[i]
            res += symb[i]

        if len(digits) > len(symb):
            res += digits[-1]

        return res
