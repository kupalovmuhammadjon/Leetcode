class Solution:
    def maxVowels(self, s: str, k: int) -> int:
        vw = "aeiou"
        res = list(s[:k])
        count = sum(1 for char in res if char in vw)

        max_count = count

        for i in range(k, len(s)):
            if res[0] in vw:
                count -= 1
            if s[i] in vw:
                count += 1
            res.pop(0)
            res.append(s[i])
            max_count = max(max_count, count)

        return max_count

# Done