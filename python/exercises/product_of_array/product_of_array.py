from typing import List


class Solution:
    def productExceptSelf(self, nums: List[int]) -> List[int]:
        result = [1 for _ in nums]
        leftProduct = 1
        rightProduct = 1
        for i in range(1, len(nums)):
            # product from items on the left
            leftProduct = leftProduct*nums[i-1]
            result[i] *= leftProduct

            # product of all items on the right
            endIndex = len(nums) - 1 - i
            rightProduct = rightProduct*nums[endIndex + 1]
            result[endIndex] *= rightProduct

        return result
