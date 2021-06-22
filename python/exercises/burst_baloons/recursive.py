import functools
from typing import List


def baloonBurstRecursive(nums: List[int]) -> int:
    nums = [1] + nums + [1]

    @functools.cache
    # i,j are indexses of first and last
    def recursive(i, j):
        if (i >= j):
            return 0
        result = 0
        for k in range(i+1, j):  # j is inclusive
            left = recursive(i, k)
            right = recursive(k, j)
            result = max(result, left + nums[i]*nums[k]*nums[j] + right)
        return result
    return recursive(0, len(nums) - 1)


def baloonBurstDP(nums: List[int]) -> int:
    nums = [1] + nums + [1]
    n = len(nums)
    dp = [[0] * n for _ in range(n)]

    def add_to_dp(i, j):
        result = 0
        for k in range(i+1, j):
            left = dp[i][k]
            right = dp[k][j]
            result = max(result, left + nums[i]*nums[k]*nums[j] + right)

        dp[i][j] = result

    for i in range(2, n):  # number of items 3,4,5
        for j in range(0, n-i):  # start index limited by out of bounds
            add_to_dp(j, j+i)
    return dp[0][n-1]


class Solution:
    def maxCoins(self, nums: List[int]) -> int:
        return baloonBurstDP(nums)
