class Solution:
    def checkIfExist(self, arr: List[int]) -> bool:
        n = len(arr)
        
        for i in range(n):
            for j in range(i + 1, n):
                if arr[j] == 2 * arr[i] or arr[i] == 2 * arr[j]:
                    return True
                elif arr[i] == 0 and arr[j] == 0:
                    return True
        return False
