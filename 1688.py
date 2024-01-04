class Solution:
    def numberOfMatches(self, n: int) -> int:

        c = 0

        while n > 0:
            if n == 1:
                break

            if n % 2:
                c += (n - 1) // 2
                n = (n - 1) // 2 + 1
 
            else:
                c += n // 2
                n //= 2

        return c

