  # Question: https://leetcode.com/problems/path-with-maximum-probability/description/ 

import heapq

class Solution:
    def maxProbability(self, n, edges, succProb, start, end):
        # Create adjacency list
        graph = [[] for _ in range(n)]
        for (u, v), prob in zip(edges, succProb):
            graph[u].append((v, prob))
            graph[v].append((u, prob))
        
        # Use a max heap to store the maximum probability path
        max_heap = [(-1.0, start)]  # Python has min heap by default, so use negative probabilities
        probabilities = [0.0] * n
        probabilities[start] = 1.0
        
        while max_heap:
            current_prob, current_node = heapq.heappop(max_heap)
            current_prob = -current_prob  # Convert back to positive
            
            if current_node == end:
                return current_prob
            
            for neighbor, edge_prob in graph[current_node]:
                new_prob = current_prob * edge_prob
                if new_prob > probabilities[neighbor]:
                    probabilities[neighbor] = new_prob
                    heapq.heappush(max_heap, (-new_prob, neighbor))
        
        return 0.0

# Example usage
n = 3
edges = [[0, 1], [1, 2], [0, 2]]
succProb = [0.5, 0.5, 0.2]
start = 0
end = 2

solution = Solution()
result = solution.maxProbability(n, edges, succProb, start, end)
print("Maximum probability from {} to {} is: {:.6f}".format(start, end, result))