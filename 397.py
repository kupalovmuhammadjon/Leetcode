class Solution:
    def integerReplacement(self, n: int) -> int:
        count = 0

        while n != 1:
            if n % 2:
                if ((n+1)//2) % 2 or ((n+1)//2) // 2 == 1:
                    n -= 1
                else:
                    n += 1
            else:
                n //= 2
            count += 1
        

        return count