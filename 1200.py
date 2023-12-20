class Solution:
    def minimumAbsDifference(self, arr: List[int]) -> List[List[int]]:
        srtd = sorted(arr)
        mn = float('inf')
        pairs = []

        for i in range(1, len(srtd)):
            diff = abs(srtd[i] - srtd[i-1])

            if diff < mn:
                mn = diff
                pairs = [[srtd[i-1], srtd[i]]]

            elif diff == mn:
                pairs.append([srtd[i-1], srtd[i]])

        return pairs

# Done passed from all tests 