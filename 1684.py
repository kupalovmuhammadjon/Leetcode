class Solution:
    def checkConsistency(self, allowed, word):
        for i in word:
            if i not in allowed:
                return False
        return True
    def countConsistentStrings(self, allowed: str, words: List[str]) -> int:
        count = 0
        for i in words:
            if self.checkConsistency(allowed, i):
                count += 1
                
        return count
        