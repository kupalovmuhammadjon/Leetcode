class Solution:
    def modifyString(self, s: str) -> str:
        s = list(s)
        
        for i in range(len(s)):
            if s[i] == "?":
                prev = s[i - 1] if i > 0 else None
                next_char = s[i + 1] if i < len(s) - 1 else None
                
                replacement = 'a'
                while replacement == prev or replacement == next_char:
                    replacement = chr(ord(replacement) + 1)

                s[i] = replacement

        return "".join(s)
