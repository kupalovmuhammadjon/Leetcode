class Solution:
    def countAndSay(self, n: int) -> str:
        if n == 1:
            return "1"

        s = "1"
        for i in range(n-1):
            s = self.counter(s)
        
        return s

    def counter(self, s):
        res = []
        s = list(s)
        i = 0
        while i < len(s):
            count = 1
            while i + 1 < len(s) and s[i] == s[i + 1]:
                i += 1
                count += 1
            res.append(str(count))
            res.append(s[i])
            i += 1

        return "".join(res)


