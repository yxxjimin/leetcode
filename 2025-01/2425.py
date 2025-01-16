"""
Runtime : 2 ms
Memory  : 24.57 MB
"""
class Solution(object):
    def xorAllNums(self, nums1, nums2):
        """
        :type nums1: List[int]
        :type nums2: List[int]
        :rtype: int
        """
        ans = 0
        
        if len(nums2) % 2 != 0:
            for num in nums1:
                ans ^= num
        
        if len(nums1) % 2 != 0:
            for num in nums2:
                ans ^= num
        
        return ans
