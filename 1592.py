class Solution:
    def reorderSpaces(self, text: str) -> str:
        count = 0
        words = ""
        word = ""
        for i in text:
            if i == " ":
                if len(word) > 0:
                    words += word + " "
                    word = ""
                count += 1
            if i.isalpha():
                word += i
                
        words += word
        words = words.split()
        if count == 0:
            return text
        
        if len(words) > 1:  
            spaces = " " * (count - (len(words)-1)*(count // (len(words)-1)))

            words = f"{count // (len(words)-1) * ' '}".join(words)
            
            words += spaces
            
        else:
            words += " " * count
            words = "".join(words)


        return words
