class Solution:
    def reversePrefix(self, word: str, ch: str) -> str:
        if ch not in word:
            return word
        i = word.find(ch)
        res = word[:i+1][::-1] + word[i+1:]

        return res