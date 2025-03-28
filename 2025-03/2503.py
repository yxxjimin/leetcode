"""
Runtime : 2920 ms
Memory  : 148.24 MB
"""
import heapq
from collections import defaultdict
from typing import List


class Solution:
    def maxPoints(self, grid: List[List[int]], queries: List[int]) -> List[int]:
        m, n = len(grid), len(grid[0])
        dydx = [(0, 1), (1, 0), (-1, 0), (0, -1)]
        visited = [[False for _ in range(n)] for _ in range(m)]

        max_level, curr_level = max(queries), 1
        cache = defaultdict(int)
        
        pq = []
        heapq.heappush(pq, (grid[0][0], (0, 0)))
        visited[0][0] = True

        while len(pq) > 0:
            lvl, (row, col) = heapq.heappop(pq)
            curr_level = max(curr_level, lvl)

            cache[curr_level] += 1

            for dy, dx in dydx:
                nr, nc = row + dy, col + dx
                if 0 <= nr < m and 0 <= nc < n and not visited[nr][nc]:
                    heapq.heappush(pq, (grid[nr][nc], (nr, nc)))
                    visited[nr][nc] = True

        for i in range(1, max_level):
            cache[i] += cache[i-1]

        return [cache[q - 1] for q in queries]
