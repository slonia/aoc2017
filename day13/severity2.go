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
	var delay int
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

	notFound := true
	for notFound {
		notFound = false
		for i := 0; i <= maxLayer; i++ {
			el, ok := firewall[i]
			if ok {
				scannerPos := pos(el, i+delay)
				if scannerPos == 0 {
					notFound = true
					break
				}
			}
		}
		delay++
	}
	fmt.Println(delay - 1)
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
