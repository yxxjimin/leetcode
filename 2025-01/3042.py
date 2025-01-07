"""
Runtime : 21 ms
Memory  : 12.53 MB
"""
class Solution(object):
    def isPrefixAndSuffix(self, str1, str2):
        return str2.startswith(str1) and str2.endswith(str1)

    def countPrefixSuffixPairs(self, words):
        """
        :type words: List[str]
        :rtype: int
        """
        cnt = 0
        for i in range(len(words) - 1):
            for j in range(i + 1, len(words)):
                if self.isPrefixAndSuffix(words[i], words[j]):
                    cnt += 1
        return cnt
    