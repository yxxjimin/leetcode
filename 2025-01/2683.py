from typing import List

"""
Runtime : 20 ms
Memory  : 22.4 MB
"""
class Solution:
    def doesValidArrayExist(self, derived: List[int]) -> bool:
        return sum(derived) % 2 == 0
