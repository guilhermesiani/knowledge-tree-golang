package algorithms

import "math"

var processed []string

func dijkstra(
	graph map[string]map[string]float64,
	costs map[string]float64,
	parents map[string]string,
) []string {
	var node string
	node = findOnLowestCost(costs)
	for node != "" {
		cost := costs[node]
		neighbors := graph[node]
		for key, _ := range neighbors {
			newCost := cost + neighbors[key]
			if costs[key] > newCost {
				costs[key] = newCost
				parents[key] = node
			}
		}
		processed = append(processed, node)
		node = findOnLowestCost(costs)
	}

	return processed
}

func findOnLowestCost(costs map[string]float64) string {
	lowestCost := math.Inf(0)
	nodeLowestCost := ""
	for node, cost := range costs {
		if cost < lowestCost && !processedHas(node) {
			lowestCost = cost
			nodeLowestCost = node
		}
	}
	return nodeLowestCost
}

func processedHas(node string) bool {
	for _, n := range processed {
		if n == node {
			return true
		}
	}

	return false
}
