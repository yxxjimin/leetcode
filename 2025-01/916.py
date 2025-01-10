"""
Runtime : 803 ms
Memory  : 15.89 MB
Comment :
    - if (A subset of B) and (B subset of C), then (A subset of C)
    - reduce `words2` into a single word
"""
class Solution(object):
    def wordSubsets(self, words1, words2):
        """
        :type words1: List[str]
        :type words2: List[str]
        :rtype: List[str]
        """
        def count(word):
            word_cnt = [0 for _ in range(26)]
            for ch in word:
                word_cnt[ord(ch) - ord("a")] += 1
            return word_cnt
        
        reduced = [0 for _ in range(26)]
        for word in words2:
            for idx, ch_cnt in enumerate(count(word)):
                reduced[idx] = max(reduced[idx], ch_cnt)

        answer = []
        for word in words1:
            if all(x >= y for x, y in zip(count(word), reduced)):
                answer.append(word)

        return answer
