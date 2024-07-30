import sys

class Solution:
    def findCheapestPrice(self, n, flights, src, dst, K):
        # Initialize the distances array with infinity
        dist = [sys.maxsize] * n
        dist[src] = 0

        # Perform Bellman-Ford algorithm for K+1 times
        for i in range(K+1):
            # Make a copy of the current distances
            temp_dist = dist[:]
            # Relax the edges
            for u, v, w in flights:
                if dist[u] != sys.maxsize and dist[u] + w < temp_dist[v]:
                    temp_dist[v] = dist[u] + w
            dist = temp_dist

        # If the distance to the destination is still infinity, return -1
        return -1 if dist[dst] == sys.maxsize else dist[dst]

# Example usage:
n = 3
flights = [[0, 1, 100], [1, 2, 100], [0, 2, 500]]
src = 0
dst = 2
K = 1

solution = Solution()
result = solution.findCheapestPrice(n, flights, src, dst, K)
print("Cheapest price from {} to {} with at most {} stops is: {}".format(src, dst, K, result))