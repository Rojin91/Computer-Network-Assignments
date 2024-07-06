import collections
import heapq
from typing import List

class Solution:
    def networkDelayTime(self, times: List[List[int]], n: int, k: int) -> int:
        # Create a graph from the given edge list
        edges = collections.defaultdict(list)
        for u, v, w in times:
            edges[u].append((v, w))
        
        # Initialize the min heap with the starting node k and time 0
        minHeap = [(0, k)]
        visit = set()
        t = 0
        
        while minHeap:
            # Pop the smallest element from the heap
            w1, n1 = heapq.heappop(minHeap)
            if n1 in visit:
                continue
            
            visit.add(n1)
            t = max(t, w1)
            
            # Push all neighboring nodes to the heap
            for n2, w2 in edges[n1]:
                if n2 not in visit:
                    heapq.heappush(minHeap, (w1 + w2, n2))
        
        # If all nodes are visited, return the total time, else return -1
        return t if len(visit) == n else -1
