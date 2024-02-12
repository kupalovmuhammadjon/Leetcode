class Solution:
    def maxProduct(self, words: List[str]) -> int:
        words = sorted(words, key=len, reverse=1)
        mx = 0
        for i in range(len(words)):
            for j in range(i + 1, len(words), 1):
                if len(set(words[i]).intersection(set(words[j]))) == 0:
                    if mx < len(words[i]) * len(words[j]):
                        mx = len(words[i]) * len(words[j])
                    break
        return mx



