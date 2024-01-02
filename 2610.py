class Solution:
    def findMatrix(self, nums: List[int]) -> List[List[int]]:
        nums.sort(reverse=True)
        result = []

        for num in nums:
            inserted = False

            for row in result:
                if num not in row:
                    row.append(num)
                    inserted = True
                    break

            if not inserted:
                result.append([num])

        return result
