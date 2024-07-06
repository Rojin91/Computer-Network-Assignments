#include <stdio.h>

#define INF 999 // Define a large value to represent infinity

void dijkstra(int n, int cost[10][10], int s, int dist[10]);

int main() {
    int i, j, n, s, cost[10][10], dist[10];
    printf("Enter the number of nodes:\n");
    scanf("%d", &n);
    printf("Enter the cost matrix:\n");
    for(i = 0; i < n; i++) {
        for(j = 0; j < n; j++) {
            scanf("%d", &cost[i][j]);
            if(cost[i][j] == 0 && i != j) { // Set 0 to INF for non-diagonal elements
                cost[i][j] = INF;
            }
        }
    }
    printf("Enter the source vertex:\n");
    scanf("%d", &s);
    dijkstra(n, cost, s, dist);
    printf("The shortest paths from %d are:\n", s);
    for(i = 0; i < n; i++) {
        if(i != s) {
            printf("%d -> %d = %d\n", s, i, dist[i]);
        }
    }
    return 0;
}

void dijkstra(int n, int cost[10][10], int s, int dist[10]) {
    int visited[10], count, min, u;

    // Initialize distances and visited array
    for(int i = 0; i < n; i++) {
        visited[i] = 0;
        dist[i] = cost[s][i];
    }
    visited[s] = 1;
    dist[s] = 0;
    count = 1;

    while(count < n - 1) {
        min = INF;
        
        // Find the vertex with the minimum distance from the set of vertices not yet processed
        for(int i = 0; i < n; i++) {
            if(dist[i] < min && !visited[i]) {
                min = dist[i];
                u = i;
            }
        }
        visited[u] = 1;
        count++;

        // Update dist value of the adjacent vertices of the picked vertex
        for(int i = 0; i < n; i++) {
            if(!visited[i] && dist[u] + cost[u][i] < dist[i]) {
                dist[i] = dist[u] + cost[u][i];
            }
        }
    }
}
