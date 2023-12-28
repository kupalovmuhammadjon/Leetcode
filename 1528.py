class Solution:
    def restoreString(self, s: str, indices: List[int]) -> str:
        n = len(s)
        result = [""] * n

        for i, index in enumerate(indices):
            result[index] = s[i]

        return "".join(result)
