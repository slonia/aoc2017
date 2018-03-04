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
	parent           string
	label            string
	weight           int
	children         []string
	calculatedWeight int
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
			nodes[label] = &Node{label: label, parent: "-", children: make([]string, 0)}
		}
		nodes[label].weight = weight
		if len(parts) > 1 {
			setParent(strings.Split(parts[1], ","), label)
		}
	}
	rootLabel := findParent(start)
	for _, node := range nodes {
		parNode, ok := nodes[node.parent]
		if ok {
			parNode.children = append(parNode.children, node.label)
		}
	}
	tree := []string{rootLabel}
	for len(tree) > 0 {
		root := nodes[tree[0]]
		tree = tree[1:]
		nodes[root.label].calculatedWeight = calcWeight(*root)
		for _, el := range root.children {
			tree = append(tree, el)
		}
	}
	tree = []string{rootLabel}
	wrongNode := nodes[rootLabel]
	for len(tree) > 0 {
		node := nodes[tree[0]]
		tree = tree[1:]
		// fmt.Println(node.label, isBalanced(*node))
		if !isBalanced(*node) {
			wrongNode = node
			for _, child := range nodes[node.label].children {
				tree = append(tree, child)
			}
		}
	}
	for _, label := range nodes[wrongNode.label].children {
		child := nodes[label]
		fmt.Println(label, child.weight, child.calculatedWeight)
	}
}

func isBalanced(node Node) bool {
	ln := len(node.children)
	for i := 0; i < ln-1; i++ {
		child1 := nodes[node.children[i]]
		child2 := nodes[node.children[i+1]]
		if child1.calculatedWeight != child2.calculatedWeight {
			return false
		}
	}
	return true
}

func calcWeight(node Node) int {
	weight := node.weight
	for _, label := range node.children {
		child := nodes[label]
		weight += child.weight
	}
	return weight
}
func findParent(start string) string {
	node := nodes[start]
	for {
		if node.parent != "-" {
			node = nodes[node.parent]
		} else {
			break
		}
	}
	return node.label
}

func setParent(parts []string, parent string) {
	for _, part := range parts {
		label := strings.Trim(part, " ")
		node, ok := nodes[label]
		if ok {
			node.parent = parent
		} else {
			nodes[label] = &Node{label: label, parent: parent, children: make([]string, 0)}
		}
	}
}
