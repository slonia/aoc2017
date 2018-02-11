package main

import (
	"bufio"
	"os"
	"strings"
	"fmt"
)

type Rule struct {
	from, to []string
}

var field []string = []string{".#.", "..#", "###"}
var rules []Rule

const iterations int = 1

func main() {
	rules = make([]Rule, 0)
	file, _ := os.Open("input.txt")
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		line := scanner.Text()
		split := strings.Split(line, " => ")
		from := strings.Split(split[0], "/")
		to := strings.Split(split[1], "/")
		rules = append(rules, Rule{from, to})
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
	for i:=0; i < curLen/step; i++ {
		square := make([]string, 0)
		for j:=0; j < step; j++ {
			square = append(square, field[j][(step*i):(step*(i+1))])
		}
		rule := findRule(square)
	}
}

func findRule(square []string) Rule {
	ln := len(square)
	for _, rule := range rules {
		for steps := 0; steps < 4; steps++ {
			square = rotate(square)
		}
	}
}

func rotate(square []string) []string {
	ln := len(square)
	newSquare = make([]string, ln)
	return square
}
