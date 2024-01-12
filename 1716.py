class Solution:
    def totalMoney(self, n: int) -> int:
        
        totalSum = 0
        saving = 1
        for i in range(n):
            if i % 7 == 0 and i != 0:
                saving = i // 7 + 1

            totalSum += saving
            saving += 1

        return totalSum