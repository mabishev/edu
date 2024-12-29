package edu

import (
	"fmt"
	"math"
)

var processed = map[string]bool{} // для отслеживания уже обработанных узлов
// алгоритм Дейкстры
func Dijkstras() {
	var graph = map[string]map[string]int{
		"start": {"a": 6, "b": 2},   // веса ребер
		"a":     {"fin": 1},         // остальные узлы и их соседи
		"b":     {"a": 3, "fin": 5}, // остальные узлы и их соседи
		"fin":   {},                 // у конечного узла нет соседей
	}

	// стоимость узла определяет, сколько времени потребуется для перехода к этому узлу
	// от начального узла
	var costs = map[string]int{
		"a":   6,
		"b":   2,
		"fin": int(math.Inf(1)),
	}

	// для вычисления итогового пути
	var parents = map[string]string{
		"a":   "start",
		"b":   "start",
		"fin": "",
	}

	node := findLowestCostNode(costs)
	for node != "" {
		cost := costs[node]
		neighbors := graph[node]
		for n := range neighbors {
			newCost := cost + neighbors[n]
			if costs[n] > newCost {
				costs[n] = newCost
				parents[n] = node
			}
		}
		processed[node] = true
		node = findLowestCostNode(costs)
	}
	fmt.Println(costs, "costs")
	fmt.Println(parents, "parents")
}

func findLowestCostNode(costs map[string]int) string {
	min := math.Inf(1)
	var lowest string
	for node := range costs {
		cost := costs[node]
		if _, ok := processed[node]; !ok {
			if cost < int(min) {
				min = float64(cost)
				lowest = node

			}
		}
	}
	return lowest
}
