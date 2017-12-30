package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// This `person` struct type has `name` and `age` fields.
type Node struct {
	parent string
	label  string
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
		_, ok := nodes[label]
		if !ok {
			nodes[label] = &Node{label: label, parent: "-"}
		}
		if len(parts) > 1 {
			setParent(strings.Split(parts[1], ","), label)
		}
	}
	node := nodes[start]
	var ok bool
	for {
		node, ok = nodes[node.parent]
		if !ok {
			break
		}
		fmt.Println(node.label)
	}
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
