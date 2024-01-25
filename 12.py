class Solution:
    def intToRoman(self, num: int) -> str:
        rooms = []
        n = num
        while n > 0:
            rooms.append(n % 10)
            n //= 10

        rooms.reverse()
        l = len(rooms)
        res = ""

        if l == 1:
            if rooms[0] == 4:
                res += "IV"
            elif rooms[0] < 4:
                res += "I" * rooms[0]
            elif rooms[0] == 5:
                res += "V"
            elif rooms[0] == 9:
                res += "IX"
            else:
                res += "V" + (rooms[0] - 5) * "I"

        elif l == 2:
            # Tens
            if rooms[0] == 4:
                res += "XL"
            elif rooms[0] < 4:
                res += "X" * rooms[0]
            elif rooms[0] == 5:
                res += "L"
            elif rooms[0] == 9:
                res += "XC"
            else:
                res += "L" + (rooms[0] - 5) * "X"

            # Ones
            if rooms[1] == 4:
                res += "IV"
            elif rooms[1] < 4:
                res += "I" * rooms[1]
            elif rooms[1] == 5:
                res += "V"
            elif rooms[1] == 9:
                res += "IX"
            else:
                res += "V" + (rooms[1] - 5) * "I"

        elif l == 3:
            # Hundreds
            if rooms[0] == 4:
                res += "CD"
            elif rooms[0] < 4:
                res += "C" * rooms[0]
            elif rooms[0] == 5:
                res += "D"
            elif rooms[0] == 9:
                res += "CM"
            else:
                res += "D" + (rooms[0] - 5) * "C"

            # Tens
            if rooms[1] == 4:
                res += "XL"
            elif rooms[1] < 4:
                res += "X" * rooms[1]
            elif rooms[1] == 5:
                res += "L"
            elif rooms[1] == 9:
                res += "XC"
            else:
                res += "L" + (rooms[1] - 5) * "X"

            # Ones
            if rooms[2] == 4:
                res += "IV"
            elif rooms[2] < 4:
                res += "I" * rooms[2]
            elif rooms[2] == 5:
                res += "V"
            elif rooms[2] == 9:
                res += "IX"
            else:
                res += "V" + (rooms[2] - 5) * "I"

        elif l == 4:
            # Thousands
            res += "M" * rooms[0]

            # Hundreds
            if rooms[1] == 4:
                res += "CD"
            elif rooms[1] < 4:
                res += "C" * rooms[1]
            elif rooms[1] == 5:
                res += "D"
            elif rooms[1] == 9:
                res += "CM"
            else:
                res += "D" + (rooms[1] - 5) * "C"

            # Tens
            if rooms[2] == 4:
                res += "XL"
            elif rooms[2] < 4:
                res += "X" * rooms[2]
            elif rooms[2] == 5:
                res += "L"
            elif rooms[2] == 9:
                res += "XC"
            else:
                res += "L" + (rooms[2] - 5) * "X"

            # Ones
            if rooms[3] == 4:
                res += "IV"
            elif rooms[3] < 4:
                res += "I" * rooms[3]
            elif rooms[3] == 5:
                res += "V"
            elif rooms[3] == 9:
                res += "IX"
            else:
                res += "V" + (rooms[3] - 5) * "I"

        return res


# Test cases
print(Solution().intToRoman(3))
print(Solution().intToRoman(8))
print(Solution().intToRoman(58))
print(Solution().intToRoman(100))
print(Solution().intToRoman(1994))
print(Solution().intToRoman(3999))
