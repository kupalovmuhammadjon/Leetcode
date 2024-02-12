class Solution:
    def lexicalOrder(self, n: int) -> List[int]:
        ls = [i for i in range(1, n+1)]

        return sorted(ls, key=lambda x: str(x))