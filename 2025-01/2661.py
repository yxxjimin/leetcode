"""
Runtime : 139 ms
Memory  : 47.17 MB
Comment :
    - Track the maximum index to complete each row and column
"""
class Solution:
    def firstCompleteIndex(self, arr: list[int], mat: list[list[int]]) -> int:
        height, width = len(mat), len(mat[0])
        num_to_idx = {arr[i]: i for i in range(len(arr))}
        answer = float("inf")

        for row in mat:
            answer = min(answer, max([num_to_idx[num] for num in row]))

        for i in range(width):
            col = [mat[j][i] for j in range(height)]
            answer = min(answer, max([num_to_idx[num] for num in col]))

        return answer