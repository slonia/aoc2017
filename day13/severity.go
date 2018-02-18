package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var firewall map[int]int

func main() {
	var severity int
	firewall = make(map[int]int)
	file, _ := os.Open("input.txt")
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	var maxLayer int
	for scanner.Scan() {
		line := scanner.Text()
		strs := strings.Split(line, ": ")
		n1, _ := strconv.Atoi(strs[0])
		n2, _ := strconv.Atoi(strs[1])
		firewall[n1] = n2
		maxLayer = n1
	}
	for i := 1; i <= maxLayer; i++ {
		el, ok := firewall[i]
		if ok {
			scannerPos := pos(el, i)
			// fmt.Printf("Iteration: %v, position: %v\n", i, scannerPos)
			if scannerPos == 0 {
				fmt.Printf("%v %v\n", i, el)
				severity += i * el
			}
		} else {
			// fmt.Printf("Iteration: %v, position: %v\n", i, 0)

		}
	}
	fmt.Println(severity)
}

func pos(deep int, step int) int {
	scannerPos := 0
	down := true
	for j := 0; j < step; j++ {
		if down {
			scannerPos++
		} else {
			scannerPos--
		}
		if scannerPos == deep-1 {
			down = false
		}
		if scannerPos == 0 {
			down = true
		}
	}
	return scannerPos
}
