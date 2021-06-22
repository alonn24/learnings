import functools
from typing import List


class Solution:
    def maxCoins(self, nums: List[int]) -> int:
        nums = [1] + nums + [1]

        @functools.cache
        def baloonBurstRecursive(i: int, j: int) -> int:
            # n = len(nums)
            if (i >= j):
                return 0
            result = 0
            for k in range(i+1, j):
                left = baloonBurstRecursive(i, k)
                right = baloonBurstRecursive(k, j)
                current = nums[i]*nums[k]*nums[j]
                result = max(result, left + right + current)
            return result

        return baloonBurstRecursive(0, len(nums) - 1)
