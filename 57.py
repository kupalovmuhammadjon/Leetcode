class Solution:
    def insert(self, intervals: List[List[int]], newInterval: List[int]) -> List[List[int]]:
        i = 0
        intervals.append(newInterval)
        intervals.sort(key=lambda x: x[0])
        while i < len(intervals):
            if i > 0 and intervals[i][0] <= intervals[i-1][1]:
                if intervals[i][1] >= intervals[i-1][0]:
                    merged = [min(intervals[i][0], intervals[i-1][0]), max(intervals[i][1], intervals[i-1][1])]
                    intervals.pop(i)
                    intervals[i-1] = merged
                else:
                    i += 1
            else:
                i += 1

        return intervals