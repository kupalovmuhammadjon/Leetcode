class Solution:
    def makeGood(self, s: str) -> str:
        s = list(s)
        l = len(s)
        i = 0

        while i < l - 1:
            if ord(s[i]) - 32 == ord(s[i + 1]) or ord(s[i]) == ord(s[i + 1]) - 32:
                s.pop(i)
                s.pop(i)
                l -= 2
                if i > 0:
                    i -= 1
            else:
                i += 1

        return "".join(s)
