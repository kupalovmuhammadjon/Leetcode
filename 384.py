class Solution:
    def __init__(self, nums: List[int]):
        self.ls = nums.copy()
        self.sh = nums.copy()

    def reset(self) -> List[int]:
        return self.ls

    def shuffle(self) -> List[int]:
        random.shuffle(self.sh)
        return self.sh

# Your Solution object will be instantiated and called as such:
# obj = Solution(nums)
# param_1 = obj.reset()
# param_2 = obj.shuffle()