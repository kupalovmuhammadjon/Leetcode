class Solution:
    def trailingZeroes(self, n: int) -> int:
        factorial = 1
        for i in range(1, n + 1):
            factorial *= i

        count = 0
        while factorial % 10 == 0:
            count += 1
            factorial //= 10

        return count
