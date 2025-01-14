"""
Runtime : 16 ms
Memory  : 12.56 MB
"""
class Solution(object):
    def findThePrefixCommonArray(self, A, B):
        """
        :type A: List[int]
        :type B: List[int]
        :rtype: List[int]
        """
        C = [0 for _ in range(len(A))]
        A_curr, B_curr = set(), set()
        
        for i in range(len(A)):
            A_curr.add(A[i])
            B_curr.add(B[i])
            C[i] = (
                C[i - 1] + int(A[i] in B_curr) + int(B[i] in A_curr)
                if A[i] != B[i]
                else C[i - 1] + 1
            )

        return C
    

"""
Runtime : 8 ms
Memory  : 12.48 MB
"""
class Solution2(object):
    def findThePrefixCommonArray(self, A, B):
        """
        :type A: List[int]
        :type B: List[int]
        :rtype: List[int]
        """
        C = [0 for _ in range(len(A))]
        freq = [0 for _ in range(len(A) + 1)]
        curr_cnt = 0
        
        for i in range(len(A)):
            freq[A[i]] += 1
            if freq[A[i]] == 2:
                curr_cnt += 1
            
            freq[B[i]] += 1
            if freq[B[i]] == 2:
                curr_cnt += 1

            C[i] = curr_cnt

        return C
