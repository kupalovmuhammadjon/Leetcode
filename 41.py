class Solution:
    def firstMissingPositive(self, nums: List[int]) -> int:
        
        l = len(nums)

        # for i in range(l):
        #     box = nums[i]
        #     j = i - 1
        #     while j >= 0 and nums[j] > box:
        #         if nums[j] > nums[j + 1]:
        #             nums[j + 1] = nums[j]
        #         j -= 1
        #         nums[j + 1] = box
        nums.sort()

        for i in range(1, l+2, 1):
            if self.binarysearch(nums, i, l) == -1:
                return i

        return l+1

    def binarysearch(self, arr, tar, l):
        start = 0
        finish = l - 1

        while start <= finish:
            mid = (start + finish) // 2

            if arr[mid] == tar:
                return tar
            elif arr[mid] > tar:
                finish = mid - 1
            else:
                start = mid + 1

        return -1

