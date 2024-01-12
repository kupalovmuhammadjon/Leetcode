class Solution:
    def thousandSeparator(self, n: int) -> str:

        if len(str(n)) <= 3:
            return str(n)

        res = []
        i = len(str(n)) - 1
        n = str(n)
        j = 0
        while i >= 0:

            res.append(n[i])

            if (len(res)-j) % 3 == 0 and len(res) != 0 and i > 0:
                res.append(".")
                j += 1

            i -= 1

        return "".join(res[::-1])
