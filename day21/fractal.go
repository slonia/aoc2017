package main

import (
	"bufio"
	"os"
	"strings"
)

var field []string = []string{".#.", "..#", "###"}
var rules map[string]string

const iterations int = 2

func main() {
	rules = make(map[string]string)
	file, _ := os.Open("input.txt")
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		line := scanner.Text()
		split := strings.Split(line, " => ")
		rules[split[0]] = split[1]
	}
	for i := 0; i < iterations; i++ {
		expand()
	}
}

func expand() {
	var step int
	curLen := len(field)
	if curLen%2 == 0 {
		step = 2
	} else {
		step = 3
	}
	var newField []string
	newLen := curLen / step * (step + 1)
	for i := 0; i < newLen; i++ {
		str := make([]string, newLen)
		newStr := strings.Join(str, "")
		newField = append(newField, newStr)
	}
}
