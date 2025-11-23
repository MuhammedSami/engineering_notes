package Algorithms

func dfsMap(start string, graph map[string][]string) []string {
	var output []string
	visited := make(map[string]bool)

	var dfs func([]string)

	dfs = func(subGraph []string) {
		for _, v := range subGraph {
			if _, exists := graph[v]; !exists {
				continue
			}

			if visited[v] {
				continue
			}

			dfs(graph[v])

			visited[v] = true

			output = append(output, v)
		}
	}

	visited[start] = true

	dfs(graph[start])

	return output
}
