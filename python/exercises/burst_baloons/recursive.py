# import functools
from typing import List

# @functools.cache


def baloonBurstRecursive(nums: List[int]) -> int:
    n = len(nums)
    if (n == 0):
        return 0
    result = 0
    for k in range(1, n-1):
        left = baloonBurstRecursive(nums[0:k+1])
        right = baloonBurstRecursive(nums[k:n])
        current = nums[0]*nums[k]*nums[n-1]
        result = max(result, left + right + current)
    return result


class Solution:
    def maxCoins(self, nums: List[int]) -> int:
        return baloonBurstRecursive([1] + nums + [1])
