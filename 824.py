class Solution:
    def toGoatLatin(self, sentence: str) -> str:
        
        s = sentence.split()
        res = ""
        
        for i in range(len(s)):
            if s[i][0].lower() not in "aeuio":
                word = s[i][1:] + s[i][:1] + "ma" + "a"*(i+1)
                res += word
                res += " "
            else:
                word = s[i] + "ma" + "a"*(i+1)
                res += word
                res += " "
            
        res = res.strip()
        return res
            
            