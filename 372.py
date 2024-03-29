class Solution:
    def superPow(self, a: int, b: List[int]) -> int:
        result = 1
        for digit in b:
            result = (pow(result, 10) * pow(a, digit)) % 1337

        return result
