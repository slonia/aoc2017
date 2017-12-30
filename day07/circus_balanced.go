package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// This `person` struct type has `name` and `age` fields.
type Node struct {
	parent string
	label  string
	weight int
}

var nodes map[string]*Node

func main() {
	nodes = make(map[string]*Node)
	file, _ := os.Open("input.txt")
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	var start string = ""
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.SplitAfter(line, " -> ")
		label_parts := strings.Split(parts[0], " ")
		label := label_parts[0]
		if start == "" {
			start = label
		}
		w_part := label_parts[1]
		weight, _ := strconv.Atoi(w_part[1 : len(w_part)-1])
		_, ok := nodes[label]
		if !ok {
			nodes[label] = &Node{label: label, parent: "-"}
		}
		nodes[label].weight = weight
		if len(parts) > 1 {
			setParent(strings.Split(parts[1], ","), label)
		}
	}
	node := findParent(start)
	fmt.Println(node.label)
}

func findParent(start string) Node {
	node := nodes[start]
	for {
		if node.parent != "-" {
			node = nodes[node.parent]
		} else {
			break
		}
	}
	return *node
}

func setParent(parts []string, parent string) {
	for _, part := range parts {
		label := strings.Trim(part, " ")
		node, ok := nodes[label]
		if ok {
			node.parent = parent
		} else {
			nodes[label] = &Node{label: label, parent: parent}
		}
	}
}
