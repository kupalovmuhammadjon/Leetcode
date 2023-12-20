class Solution:
    def removeStars(self, s: str) -> str:
        ls = list(s)
        result = []

        for char in ls:
            if char == '*':
                if result:
                    result.pop()  
            else:   
                result.append(char)

        return "".join(result)

# Done