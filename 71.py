class Solution:
    def simplifyPath(self, path: str) -> str:
        path = list(path)
        i = 0
        while i < len(path):
            if i and path[i] == "/":
                if path[i-1] == "/":
                    path.pop(i)
                else:
                    i += 1
            else:
                i += 1
        if len(path) > 1 and path[-1] == "/":
            path.pop()

        ch = "".join(path).split("/")
        i = 0
        while i < len(ch):
            if i and ch[i] == "..":
                ch.pop(i-1)
                ch.pop(i-1)
                i -= 1
            elif ch[i] == "..":
                ch.pop(i)
            elif ch[i] == ".":
                ch.pop(i)
            else:
                i += 1
        

        if len(ch) == 1 and ch[0] == "":
            return "/" 
        elif len(ch) == 0:
            return "/"

        if ch and ch[0] != "":
            ch.insert(0, "")
        
        return "/".join(ch)

