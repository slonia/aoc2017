package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var edges []Edge
var maxNode int

const INF = 1000000

type Edge struct {
	a, b int
	cost int
}

func main() {
	edges = make([]Edge, 0)
	file, _ := os.Open("input.txt")
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		line := scanner.Text()
		lineNodes := strings.Split(line, "/")
		node1, _ := strconv.Atoi(lineNodes[0])
		node2, _ := strconv.Atoi(lineNodes[1])
		if node1 > maxNode {
			maxNode = node1
		}
		if node2 > maxNode {
			maxNode = node2
		}
		cost := node1 + node2
		edges = append(edges, Edge{node1, node2, cost})
	}
	dist := make([]int, maxNode+1)
	for i, _ := range dist {
		if i == 0 {
			dist[i] = 0
		} else {
			dist[i] = INF
		}
	}
	n := maxNode
	fmt.Println(edges)
	for i := 0; i < n-1; i++ {
		for _, e := range edges {
			if dist[e.a] <= INF {
				// fmt.Printf("For %v, Comparing %v+%v and %v\n", e.b, dist[e.a], e.cost, dist[e.b])
				if dist[e.a]+e.cost <= dist[e.b] {
					dist[e.b] = dist[e.a] + e.cost
				}
			}
		}
	}
	for i, el := range dist {
		fmt.Println(i, el)

	}
}
