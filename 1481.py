from collections import Counter

class Solution:
    def findLeastNumOfUniqueInts(self, arr: List[int], k: int) -> int:
        counter = Counter(arr)
        sorted_counts = sorted(counter.values())
        print(sorted_counts)
        count_removed = 0
        for count in sorted_counts:
            if k >= count:
                k -= count
                count_removed += 1
            else:
                break
        
        return len(sorted_counts) - count_removed