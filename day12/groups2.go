package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	nodes := make(map[int][]int)
	node_names := []int{}
	line_nodes := []int{}
	processed := make([]int, 0)
	groups_count := 0
	file, _ := os.Open("input.txt")
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		line_nodes = nil
		line := scanner.Text()
		line_split := strings.Split(line, " <-> ")
		node, _ := strconv.Atoi(line_split[0])
		line_nodes = append(line_nodes, node)
		line_split = strings.Split(line_split[1], ", ")
		for _, el := range line_split {
			node, _ := strconv.Atoi(el)
			line_nodes = append(line_nodes, node)
		}
		for _, el1 := range line_nodes {
			_, ok := nodes[el1]
			if !ok {
				nodes[el1] = []int{}
			}
			for _, el2 := range line_nodes {
				if !contains(nodes[el1], el2) {
					nodes[el1] = append(nodes[el1], el2)
				}
			}
		}
		curNode := node_names[0]
		for len(processed) < len(node_names) {

		}
		_, ok := nodes[0]
		if ok {
			for _, el := range nodes[0] {
				for _, to_add := range nodes[el] {
					if !contains(nodes[0], to_add) {
						nodes[0] = append(nodes[0], to_add)
					}
				}
			}
		}
	}
	fmt.Println(len(nodes[0]))
}

func contains(arr []int, el int) bool {
	for _, arr_el := range arr {
		if el == arr_el {
			return true
		}
	}
	return false
}
