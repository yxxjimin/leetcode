"""
Runtime : 274 ms
Comment :
    - DFS (within same cost) + BFS (for finding minimum costs)
"""
class Solution:
    _queue = []
    _grid = []
    _min_cost = []
    _dirs = [(0, 1), (0, -1), (1, 0), (-1, 0)]

    def minCost(self, grid: list[list[int]]) -> int:
        self._grid = grid
        self._min_cost = [
            [float("inf") for _ in range(len(grid[0]))] 
            for _ in range(len(grid))
        ]

        cur_cost = 0
        self._dfs(0, 0, cur_cost)

        while self._queue:
            cur_cost += 1
            for _ in range(len(self._queue)):
                row, col = self._queue.pop(0)
                for dy, dx in self._dirs:
                    self._dfs(row + dy, col + dx, cur_cost)

        return self._min_cost[-1][-1]

    def _visited(self, row: int, col: int) -> bool:
        return (
            row not in range(len(self._grid))
            or col not in range(len(self._grid[0]))
            or self._min_cost[row][col] < float("inf")
        )

    def _dfs(self, row: int, col: int, cost: int):
        while not self._visited(row, col):
            self._queue.append((row, col))
            self._min_cost[row][col] = cost
            dy, dx = self._dirs[self._grid[row][col] - 1]
            row += dy
            col += dx
