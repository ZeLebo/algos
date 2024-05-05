from typing import List

class Solution:
    def __init__(self):
        ...

    def productExceptSelf(self, nums: List[int]) -> List[int]:
        left = [0 for _ in range(len(nums))]
        right = [0 for _ in range(len(nums))]
        right_product = 1
    
        for index, num in enumerate(nums):
            left[index] = left[index - 1] * num if index != 0 else num

        for index, num in reversed(list(enumerate(nums))):
            right[index] = right[index + 1] * num if index != len(nums) - 1 else num

        # res = [0 for _ in range(len(nums))]
        # for index in range(len(nums)):
        #     if index == 0:
        #         res[index] = right[1]
        #     elif index == len(nums) - 1:
        #         res[index] = left[index - 1]
        #     else:
        #         res[index] = left[index - 1] * right[index + 1]

        res = [right[1] if index == 0 else left[index - 1] * right[index + 1] if index != len(nums) - 1 else left[index - 1] for index in range(len(nums))]

        return res


nums = [1, 2, 3, 4]
print(f"{Solution().productExceptSelf(nums)}")
