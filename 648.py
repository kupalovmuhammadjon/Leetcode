class Solution:
    def replaceWords(self, dictionary: List[str], sentence: str) -> str:
        words = sentence.split()
        for i in range(len(words)):
            for j in dictionary:
                if words[i].startswith(j):
                    words[i] = j

        return " ".join(words)
