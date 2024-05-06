class Solution(object):
    def longestConsecutive(self, nums):
        if len(nums) == 0:
            return 0

        nums_set = set(nums)
        max_answer = 1

        for num in nums_set:
            if num - 1 not in nums_set:
                current_num = num
                current_streak = 1
                while current_num + 1 in nums_set:
                    current_num += 1
                    current_streak += 1
                max_answer = max(max_answer, current_streak)
        return max_answer
