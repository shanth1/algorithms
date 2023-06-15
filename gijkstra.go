package algorithms

import (
	"math"
	"strings"
)

func findLowestCostNode(costs map[string]int, precessed map[string]bool) (lowestCostNode string) {
	lowestCost := math.MaxInt
	for key, value := range costs {
		cost := value
		if _, ok := precessed[key]; !ok {
			if cost < lowestCost {

				lowestCost = cost
				lowestCostNode = key
			}
		}
	}
	return
}

func getRoute(parents map[string]string) string {
	parent := parents["finish"]
	var route []string
	route = append(route, "finish")
	for parent != "start" {
		route = append(route, parent)
		parent = parents[parent]
	}
	route = append(route, "start")
	for i, j := 0, len(route)-1; i < j; i, j = i+1, j-1 {
		route[i], route[j] = route[j], route[i]
	}

	return strings.Join(route, " -> ")
}

func Dijkstra(graph map[string]map[string]int, costs map[string]int, parents map[string]string) string {
	processed := make(map[string]bool)
	node := findLowestCostNode(costs, processed)
	for node != "" {
		cost := costs[node]
		neighbors := graph[node]
		for key := range neighbors {
			newCost := cost + neighbors[key]
			if costs[key] > newCost {
				costs[key] = newCost
				parents[key] = node
			}
		}
		processed[node] = true
		node = findLowestCostNode(costs, processed)
	}

	return getRoute(parents)
}
