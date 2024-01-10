from collections import Counter

class Solution:
    def frequencySort(self, nums: List[int]) -> List[int]:
        # Count the frequency of each number using Counter
        frequency_count = Counter(nums)

        # Sort the numbers based on increasing frequency and decreasing value
        sorted_nums = sorted(nums, key=lambda x: (frequency_count[x], -x))

        return sorted_nums
