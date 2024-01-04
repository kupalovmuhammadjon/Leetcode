class Solution:
    def interpret(self, command: str) -> str:
        res = ""

        for i in range(len(command)):
            if command[i] == "(" and command[i+1] == ")":
                res += "o"
            elif command[i].isalpha():
                res += command[i]
        
        return res