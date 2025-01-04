// алгоритм Дейкстры
package main

import (
	"fmt"
	"math"
)

func main() {
	graph := make(map[string]map[string]int)
	graph["SD"] = map[string]int{}
	graph["SD"]["LA"] = 1
	graph["SD"]["NY"] = 7

	graph["LA"] = map[string]int{}
	graph["LA"]["SD"] = 1
	graph["LA"]["NY"] = 4
	graph["LA"]["LDN"] = 10

	graph["NY"] = map[string]int{}
	graph["NY"]["SD"] = 7
	graph["NY"]["LA"] = 4
	graph["NY"]["LDN"] = 3

	graph["LDN"] = map[string]int{}
	graph["LDN"]["LA"] = 10
	graph["LDN"]["NY"] = 3

	costs, parent := findShortestPath(graph, "SD", "LDN")
	fmt.Println(costs, parent)
}

func findShortestPath(graph map[string]map[string]int, startNode string, finishNode string) (map[string]int, map[string]string) {
	costs := make(map[string]int)
	costs[finishNode] = math.MaxInt
	parents := make(map[string]string)
	parents[finishNode] = ""

	//определяет путешествовали мы по этому маршруту или нет
	processed := make(map[string]bool)
	for node, cost := range graph[startNode] {
		costs[node] = cost        //стоимость от узла node равна базовой стоимости
		parents[node] = startNode //родителем этого узла будет startNode
	}
	lowestCostNode := findLowerCostNode(costs, processed)
	for lowestCostNode != "" {
		for node, cost := range graph[lowestCostNode] {
			newCost := costs[lowestCostNode] + cost
			if newCost < costs[node] {
				costs[node] = newCost
				parents[node] = lowestCostNode
			}
		}
		processed[lowestCostNode] = true
		lowestCostNode = findLowerCostNode(costs, processed)

	}
	return costs, parents
}

// расчет минимальной стоимости
func findLowerCostNode(costs map[string]int, processed map[string]bool) string {
	lowerCost := math.MaxInt
	lowestCostNode := ""

	for node, cost := range costs {
		if _, inProcessed := processed[node]; cost < lowerCost && !inProcessed {
			lowerCost = cost
			lowestCostNode = node
		}
	}
	return lowestCostNode
}
